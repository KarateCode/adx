package adx

import (
	"text/template"
	"bytes"
	"encoding/xml"
	// "io"
	// "os"
	"net/http"
)

type servicedAccountService struct {
	conn *Conn
}

type ServicedAccountGetSelector struct {
	XMLName   xml.Name `xml:"selector"`
	// XsiType   string `xml:"xsi:type,attr"`
	CustomerIds []int `xml:"customerIds"`
	EnablePaging bool `xml:"enablePaging"`
	SubmanagersOnly bool `xml:"submanagersOnly"`
}

var mccLayout = `{{define "T"}}<?xml version="1.0" encoding="UTF-8"?>
	<env:Envelope xmlns:xsd="http://www.w3.org/2001/XMLSchema" 
		xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" 
		xmlns:wsdl="https://adwords.google.com/api/adwords/mcm/{{.Auth.Version}}" 
		xmlns:env="http://schemas.xmlsoap.org/soap/envelope/">
	<env:Header>
		<wsdl:RequestHeader xmlns="https://adwords.google.com/api/adwords/cm/{{.Auth.Version}}">
			<userAgent>AwApi-Ruby-0.4.3|central</userAgent>
			<developerToken>{{.Auth.DeveloperToken}}</developerToken>
			<authToken>{{.AuthToken}}</authToken>
		</wsdl:RequestHeader>
	</env:Header>
	<env:Body>
		<get xmlns="https://adwords.google.com/api/adwords/mcm/{{.Auth.Version}}">{{.Body}}
		</get>
	</env:Body>
</env:Envelope>{{end}}`

type Account struct {
	CustomerId int64 `xml:"customerId"`
	Login string `xml:"login"`
	CompanyName string `xml:"companyName"`
	CanManageClients bool `xml:"canManageClients"`
}

type Link struct {
	ManagerId int64 `xml:"managerId>id"`
	ClientId int64 `xml:"clientId>id"`
	TypeOfLink string `xml:"typeOfLink"`
	DescriptiveName string `xml:"descriptiveName"`
	// LinkType bool `xml:"Link.Type"`
}

type ServicedAccountServiceGet struct {
	XMLName   xml.Name `xml:"Envelope"`
	Body struct {
		XMLName   xml.Name `xml:"Body"`
		GetResponse struct {
			XMLName   xml.Name `xml:"getResponse"`
			Rval struct {
				XMLName   xml.Name `xml:"rval"`
				Accounts []Account `xml:"accounts"`
				Links []Link `xml:"links"`
			}
		}
	}
}

func (self *servicedAccountService) Get(v ServicedAccountGetSelector) (*[]Account, *[]Link, error) {
	sasGet := new(ServicedAccountServiceGet)
		
	tmp, err := template.New("temp").Parse(mccLayout)
	if err != nil {
		return nil, nil, err
	}
	
	p, err := xml.MarshalIndent(v, "			", "	")
	if err != nil {
		return nil, nil, err
	}
	
	buffer := bytes.NewBufferString("")
	execErr := tmp.ExecuteTemplate(buffer, "T", data{Auth:&self.conn.Auth, AuthToken:self.conn.token, Body:string(p), Mcc:"mcm", Operation:"get"})
	if execErr != nil {
		return nil, nil, execErr
	}

	// io.Copy(os.Stdout, buffer)
	// return nil, nil
	
	// service doesn't exist in v201206
	req, err := http.NewRequest("POST", 
		"https://adwords" + self.conn.sandboxUrl + ".google.com/api/adwords/mcm/v201109/ServicedAccountService", 
		buffer)
	if err != nil {
		return nil, nil, err
	}
	
	req.Header.Add("Content-Type", "application/soap+xml") // VERY IMPORTANT. ADX wouldn't accept xml without it
	req.Header.Add("Authorization", "GoogleLogin auth=" + self.conn.token)
	req.Header.Add("clientCustomerId", self.conn.Auth.ClientId)
	req.Header.Add("developerToken", self.conn.Auth.DeveloperToken)
	
	res, err := http.DefaultClient.Do(req)  
	if err != nil {
		return nil, nil, err
	}
	defer res.Body.Close()
	
	// io.Copy(os.Stdout, res.Body) // uncomment this to view http response. Found a 414 once
	// return sasGet, nil
	
	decoder := xml.NewDecoder(res.Body)
	err = decoder.Decode(sasGet)
	if err != nil {
		panic(err)
		return nil, nil, err
	}
	
	return &sasGet.Body.GetResponse.Rval.Accounts, &sasGet.Body.GetResponse.Rval.Links, nil
}

