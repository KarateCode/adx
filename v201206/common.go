package v201206

import (
	"strings"
	"bufio"
	"net/url"
	"io"
	"os"
	"net/http"
	"text/template"
	"bytes"
	"encoding/xml"
)

type Auth struct {
	DeveloperToken, ClientId, Email, Password, Version, Type string
	Sandbox bool
}

type ApiData struct{
	Auth *Auth
	AuthToken string
	Body, Mcc, Operation string
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
<env:Envelope 
	xmlns:xsd="http://www.w3.org/2001/XMLSchema"
	xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:wsdl="https://adwords.google.com/api/adwords/cm/{{.Auth.Version}}" xmlns:env="http://schemas.xmlsoap.org/soap/envelope/">
	<env:Header>
		<wsdl:RequestHeader xmlns="https://adwords.google.com/api/adwords/cm/{{.Auth.Version}}">
			<userAgent>golang|adx</userAgent>
			<developerToken>{{.Auth.DeveloperToken}}</developerToken>
			<clientCustomerId>{{.Auth.ClientId}}</clientCustomerId>
			<authToken>{{.AuthToken}}</authToken>
		</wsdl:RequestHeader>
	</env:Header>
	<env:Body>
		<{{.Operation}} xmlns="https://adwords.google.com/api/adwords/{{.Mcc}}/{{.Auth.Version}}">{{.Body}}</{{.Operation}}>
	</env:Body>
</env:Envelope>{{end}}`

type Conn struct {
	Auth
	SandboxUrl string
	Token string
}

type Fault struct {
	XMLName   xml.Name
	FaultCode string `xml:"faultcode"`
	FaultString string `xml:"faultstring"`
}

type MutateResponse struct {
	XMLName   xml.Name `xml:"Envelope"`
	Body struct {
		Fault Fault
	}
}

type Ordering struct {
	Field string `xml:"ordering>field"`
	SortOrder string `xml:"ordering>sortOrder,omitempty"`
}

type Predicate struct {
	Field string `xml:"field"`
	Operator string `xml:"operator"`
	Values []string `xml:"values"`
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
	p, err := xml.Marshal(v)
	// p, err := xml.MarshalIndent(v, "", "	")
	if err != nil {
		return nil, err
	}
	
	// println(string(p) + "\n")
	
	buffer := bytes.NewBufferString("")
	execErr := layout.ExecuteTemplate(buffer, "T", ApiData{
		Auth:       &conn.Auth,
		AuthToken:  conn.Token,
		Body:       string(p),
		// Body:       body,
		Mcc:        "cm",
		Operation:  operation,
	})
	if execErr != nil {
		return nil, err
	}

	// io.Copy(os.Stdout, buffer)
	
	// println("https://adwords" + conn.SandboxUrl + ".google.com/api/adwords/cm/" + conn.Version + "/" + service)
	req, err := http.NewRequest("POST", "https://adwords" + conn.SandboxUrl + ".google.com/api/adwords/cm/" + conn.Version + "/" + service, buffer)
	if err != nil {
		return nil, err
	}
	
	req.Header.Add("Content-Type", "application/soap+xml") // VERY IMPORTANT. ADX wouldn't accept xml without it
	req.Header.Add("Authorization", "GoogleLogin auth=" + conn.Token)
	req.Header.Add("clientCustomerId", conn.Auth.ClientId)
	req.Header.Add("developerToken", conn.Auth.DeveloperToken)
	
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	return res.Body, nil
}
