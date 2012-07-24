package adx

import (
	"strings"
	"bufio"
	"net/url"
	"os"
	"net/http"
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
}

var layout = `{{define "T"}}<?xml version="1.0" encoding="UTF-8"?>
<soap:Envelope
  xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/"
  xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" 
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
	token string
}

func New(auth Auth) (*Conn) {
	conn := Conn{Auth:auth, token:Authenticate(auth.Email, auth.Password)}
	conn.CampaignService.conn = &conn
	conn.ServicedAccountService.conn = &conn
	
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
