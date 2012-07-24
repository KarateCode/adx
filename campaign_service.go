package adx

import (
	"text/template"
	"bytes"
	"encoding/xml"
	"io"
	"os"
	"net/http"
)

type campaignService struct {
	conn *Conn
}

type CampaignGetSelector struct {
	XMLName   xml.Name `xml:"serviceSelector"`
	// XsiType   string `xml:"xsi:type,attr"`
	Fields []string `xml:"fields"`
	Field string `xml:"predicates>field"`
	Operator string `xml:"predicates>operator"`
	Values []string `xml:"predicates>values"`
}

type data struct{
	Auth *Auth
	AuthToken string
	Body, Mcc, Operation string
}

type CampaignGet struct {
	XMLName   xml.Name `xml:"Envelope"`
	Body struct {
		XMLName   xml.Name `xml:"Body"`
		GetResponse struct {
			XMLName   xml.Name `xml:"getResponse"`
			Rval struct {
				XMLName   xml.Name `xml:"rval"`
				TotalNumEntries int `xml:"totalNumEntries"`
				TotalBudget struct {
					Period string `xml:"period"`
					Amount struct {
						MicroAmount int64 `xml:"microAmount"`
					} `xml:"amount"`
				}
				Entries []struct {
					Id int64 `xml:"id"`
					Name string `xml:"name"`
					CampaignStats struct {
						network string `xml:"network"`
						StatsType string `xml:"Stats.Type"`
					} `xml:"campaignStats"`
					FrequencyCap struct {
						Impressions int64 `xml:"impressions"`
					} `xml:"frequencyCap"`
				} `xml:"entries"`
			}
		}
	}
}

func (self *campaignService) Get(v CampaignGetSelector) (*CampaignGet, error) {
	campaignGet := new(CampaignGet)
	
	tmp, err := template.New("temp").Parse(layout)
	if err != nil {
		return nil, err
	}
	
	p, err := xml.MarshalIndent(v, "", "	")
	if err != nil {
		return nil, err
	}
	
	
	buffer := bytes.NewBufferString("")
	execErr := tmp.ExecuteTemplate(buffer, "T", data{Auth:&self.conn.Auth, AuthToken:self.conn.token, Body:string(p), Mcc:"cm", Operation:"get"})
	if execErr != nil {
		return nil, err
	}

	// io.Copy(os.Stdout, buffer)
	// return nil, nil
	
	req, err := http.NewRequest("POST", 
		"https://adwords" + self.conn.sandboxUrl + ".google.com/api/adwords/cm/" + self.conn.Version + "/CampaignService", 
		buffer)
	if err != nil {
		return nil, err
	}
	
	req.Header.Add("Content-Type", "application/soap+xml") // VERY IMPORTANT. ADX wouldn't accept xml without it
	req.Header.Add("Authorization", "GoogleLogin auth=" + self.conn.token)
	req.Header.Add("clientCustomerId", self.conn.Auth.ClientId)
	req.Header.Add("developerToken", self.conn.Auth.DeveloperToken)
	
	res, err := http.DefaultClient.Do(req)  
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	
	decoder := xml.NewDecoder(res.Body)
	err = decoder.Decode(campaignGet)
	if err != nil {
		panic(err)
		return nil, err
	}
	
	io.Copy(os.Stdout, res.Body) // uncomment this to view http response. Found a 414 once
	return campaignGet, nil
}

type BiddingStrategy struct {
	// XMLName   xml.Name `xml:"operations"`
	XsiType   string `xml:"xsi:type,attr"`
	Cm string `xml:"xmlns:cm,attr"` 
	Xsi string `xml:"xmlns:xsi,attr"`
}

type Settings struct {
	XsiType   string `xml:"xsi:type,attr"`
	OptIn bool `xml:"optIn"`
}

type CampaignMutateOperations struct {
	XMLName                 xml.Name `xml:"operations"`
	Operator                string `xml:"operator"`
	Name                    string `xml:"operand>name"`
	Status                  string `xml:"operand>status"`
	Period                  string `xml:"operand>budget>period"`
	MicroAmount             string `xml:"operand>budget>amount>microAmount"`
	DeliveryMethod          string `xml:"operand>budget>deliveryMethod"`
	BiddingStrategy         BiddingStrategy `xml:"operand>biddingStrategy"`
	Settings Settings       `xml:"operand>settings"`
	TargetGoogleSearch      bool `xml:"operand>networkSetting>targetGoogleSearch"`
	TargetSearchNetwork     bool `xml:"operand>networkSetting>targetSearchNetwork"`
	TargetContentNetwork    bool `xml:"operand>networkSetting>targetContentNetwork"`
	TargetContentContextual bool `xml:"operand>networkSetting>targetContentContextual"`
}

func (self *campaignService) Mutate(v CampaignMutateOperations) {
	v.BiddingStrategy.Cm = "https://adwords.google.com/api/adwords/cm/" + self.conn.Version
	v.BiddingStrategy.Xsi = "http://www.w3.org/2001/XMLSchema-instance"
	
	tmp, err := template.New("temp").Parse(layout)
	if err != nil {
		panic(err)
	}
	
	// v := servicedAccountServiceGet{EnablePaging:false, SubmanagersOnly:false}
	p, err := xml.MarshalIndent(v, "			", "	")
	if err != nil {
		panic(err)
	}
	
	
	buffer := bytes.NewBufferString("")
	execErr := tmp.ExecuteTemplate(buffer, "T", data{Auth:&self.conn.Auth, AuthToken:self.conn.token, Body:string(p), Mcc:"cm", Operation:"mutate"})
	if execErr != nil {
		panic(execErr)
	}

	// io.Copy(os.Stdout, buffer)
	// return
	
	println("https://adwords" + self.conn.sandboxUrl + ".google.com/api/adwords/cm/" + self.conn.Version + "/CampaignService")
	req, err := http.NewRequest("POST", 
		"https://adwords" + self.conn.sandboxUrl + ".google.com/api/adwords/cm/" + self.conn.Version + "/CampaignService", 
		buffer)
	if err != nil {
		panic(err)
	}
	
	req.Header.Add("Content-Type", "application/soap+xml") // VERY IMPORTANT. ADX wouldn't accept xml without it
	req.Header.Add("Authorization", "GoogleLogin auth=" + self.conn.token)
	req.Header.Add("clientCustomerId", self.conn.Auth.ClientId)
	req.Header.Add("developerToken", self.conn.Auth.DeveloperToken)
	
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	
	io.Copy(os.Stdout, res.Body) // uncomment this to view http response. Found a 414 once
	return 
}

