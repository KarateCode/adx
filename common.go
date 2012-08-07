package adx

import (
	"strings"
	"bufio"
	"net/url"
	"os"
	"net/http"
	"text/template"
	"bytes"
	"io"
	// "os"
	"encoding/xml"
)

type Auth struct {
	DeveloperToken, ClientId, Email, Password, Version, Type string
	Sandbox bool
}

var AdxPush Auth
var AdxPull Auth

func init() {
	sandbox := false
	if os.Getenv("AdxPushSandbox") == "true" {
		sandbox = true
	}
	
	AdxPush = Auth{
		DeveloperToken: os.Getenv("AdxPushDeveloperToken"), 
		ClientId:       os.Getenv("AdxPushClientId"), 
		Type:           os.Getenv("AdxPushType"), 
		Email:          os.Getenv("AdxPushEmail"),
		Password:       os.Getenv("AdxPushPassword"), 
		Version:        os.Getenv("AdxPushVersion"), 
		Sandbox:        sandbox,
	}
	
	sandbox = false
	if os.Getenv("AdxPullSandbox") == "true" {
		sandbox = true
	}
	
	AdxPull = Auth{
		DeveloperToken: os.Getenv("AdxPullDeveloperToken"), 
		ClientId:       os.Getenv("AdxPullClientId"), 
		Type:           os.Getenv("AdxPullType"), 
		Email:          os.Getenv("AdxPullEmail"),
		Password:       os.Getenv("AdxPullPassword"), 
		Version:        os.Getenv("AdxPullVersion"), 
		Sandbox:        sandbox,
	}
	
	var err error
	layout, err = template.New("temp").Parse(layoutString)
	if err != nil {
		panic(err)
	}
}

var layout *template.Template
var layoutString = `{{define "T"}}<?xml version="1.0" encoding="UTF-8"?>
<soap:Envelope
  xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/"
  xmlns="https://adwords.google.com/api/adwords/{{.Mcc}}/{{.Auth.Version}}">
	<soap:Header>
		<RequestHeader>
			<userAgent>golang|adx</userAgent>
			<developerToken>{{.Auth.DeveloperToken}}</developerToken>
			<clientCustomerId>{{.Auth.ClientId}}</clientCustomerId>
			<authToken>{{.AuthToken}}</authToken>
		</RequestHeader>
	</soap:Header>
	<soap:Body>
		<{{.Operation}}>{{.Body}}
		</{{.Operation}}>
	</soap:Body>
</soap:Envelope>{{end}}`

type Conn struct {
	Auth
	sandboxUrl string
	CampaignService campaignService
	ServicedAccountService servicedAccountService
	AdgroupCriterionService adgroupCriterionService
	AdgroupService adgroupService
	token string
}

func New(auth Auth) (*Conn) {
	conn := Conn{Auth:auth, token:Authenticate(auth.Email, auth.Password)}
	conn.CampaignService.conn = &conn
	conn.ServicedAccountService.conn = &conn
	conn.AdgroupCriterionService.conn = &conn
	conn.AdgroupService.conn = &conn
	
	if auth.Sandbox {
		conn.sandboxUrl = "-sandbox"
	}
	
	return &conn
}

func Authenticate(email, password string) string {
	var authToken string
	client := new(http.Client)
	
	var params url.Values = make(url.Values)
	params.Add("Email", email)
	params.Add("Passwd", password)
	params.Add("accountType", "GOOGLE")
	params.Add("source", "adwords-tutorial")
	params.Add("service", "adwords")
	
	res, err := client.PostForm("https://www.google.com/accounts/ClientLogin", params)
	if err != nil {
		panic(err)
	}
	// io.Copy(os.Stdout, res.Body)
	lines := bufio.NewReader(res.Body)
	for {
		line, _, err := lines.ReadLine()
		if err != nil {
			// return ""// seems to take this one as the exit
			break
		} else if len(line) < 0 {
			println("error reading")
			os.Exit(1)
		} else if len(line) == 0 { // EOF
			println("reading: nr == zero")
		} else if len(line) > 0 {
			text := string(line[:])
			if (strings.Contains(text, "Auth=")) {
				authToken = strings.Replace(text, "Auth=", "", -1)
				break
			}
		}
	}
	return authToken
}

func CallApi(v interface{}, conn *Conn, service string, operation string) (io.ReadCloser, error) {
	p, err := xml.MarshalIndent(v, "", "	")
	if err != nil {
		return nil, err
	}
	
	buffer := bytes.NewBufferString("")
	execErr := layout.ExecuteTemplate(buffer, "T", data{
		Auth:       &conn.Auth, 
		AuthToken:  conn.token, 
		Body:       string(p), 
		Mcc:        "cm", 
		Operation:  operation,
	})
	if execErr != nil {
		return nil, err
	}

	// io.Copy(os.Stdout, buffer)
	// return nil, nil
	
	req, err := http.NewRequest("POST", 
		"https://adwords" + conn.sandboxUrl + ".google.com/api/adwords/cm/" + conn.Version + "/" + service, 
		buffer)
	if err != nil {
		return nil, err
	}
	
	req.Header.Add("Content-Type", "application/soap+xml") // VERY IMPORTANT. ADX wouldn't accept xml without it
	req.Header.Add("Authorization", "GoogleLogin auth=" + conn.token)
	req.Header.Add("clientCustomerId", conn.Auth.ClientId)
	req.Header.Add("developerToken", conn.Auth.DeveloperToken)
	
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	return res.Body, nil
}