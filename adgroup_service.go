package adx

import (
	// "fmt"
	"text/template"
	"bytes"
	"encoding/xml"
	"io"
	"os"
	"net/http"
)

type adgroupService struct {
	conn *Conn
}

type AdgroupPredicate struct {
	Field string `xml:"field"`
	Operator string `xml:"operator"`
	Values []string `xml:"values"`
}

type AdgroupSelector struct {
	XMLName   xml.Name `xml:"serviceSelector"`
	// XsiType   string `xml:"xsi:type,attr"`
	Fields []string `xml:"fields"`
	// Field string `xml:"predicates>field"`
	// Operator string `xml:"predicates>operator"`
	// Values []string `xml:"predicates>values"`
	Predicates []Predicate `xml:"predicates"`
	StartIndex int `xml:"paging>startIndex"`
	NumberResults int `xml:"paging>numberResults"`
}

type AdgroupGet struct {
	XMLName   xml.Name `xml:"Envelope"`
	Body struct {
		// XMLName   xml.Name
		GetResponse struct {
			// XMLName   xml.Name `xml:"getResponse"`
			Rval struct {
				// XMLName   xml.Name
				TotalNumEntries int `xml:"totalNumEntries"`
				Entries []struct {
					Id int64 `xml:"id"`
					Name string `xml:"name"`
					Status string `xml:"status"`
					Bids struct {
						AdGroupBidsType string `xml:"AdGroupBids.Type"`	
						EnhancedCpcEnabled bool `xml:"enhancedCpcEnabled"`
					} `xml:"bids"`
					Stats struct {
						Network string `xml:"network"`
						StatsType string `xml:"Stats.Type"`
					} `xml:"stats"`
				} `xml:"entries"`
			} `xml:"rval"`
		} `xml:"getResponse"`
	}
}

func (self *adgroupService) Get(v AdgroupSelector) (*AdgroupGet, error) {
	adgroupGet := new(AdgroupGet)
	
	tmp, err := template.New("temp").Parse(layout)
	if err != nil {
		return nil, err
	}
	
	p, err := xml.MarshalIndent(v, "", "	")
	if err != nil {
		return nil, err
	}
	
	buffer := bytes.NewBufferString("")
	execErr := tmp.ExecuteTemplate(buffer, "T", data{
		Auth:&self.conn.Auth, 
		AuthToken:self.conn.token, 
		Body:string(p), 
		Mcc:"cm", 
		Operation:"get",
	})
	if execErr != nil {
		return nil, err
	}

	// io.Copy(os.Stdout, buffer)
	// return nil, nil
	
	req, err := http.NewRequest("POST", 
		"https://adwords" + self.conn.sandboxUrl + ".google.com/api/adwords/cm/" + self.conn.Version + "/AdGroupService", 
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
	
	// io.Copy(os.Stdout, res.Body) // uncomment this to view http response. Found a 414 once
	// return adgroupGet, nil
	
	decoder := xml.NewDecoder(res.Body)
	err = decoder.Decode(adgroupGet)
	if err != nil {
		panic(err)
		return nil, err
	}
	
	// fmt.Printf("\nadgroupGet%v\n", adgroupGet)
	// io.Copy(os.Stdout, res.Body) // uncomment this to view http response. Found a 414 once
	return adgroupGet, nil
}

// type BiddingStrategy struct {
// 	// XMLName   xml.Name `xml:"operations"`
// 	XsiType   string `xml:"xsi:type,attr"`
// 	Cm string `xml:"xmlns:cm,attr"` 
// 	Xsi string `xml:"xmlns:xsi,attr"`
// }

// type Settings struct {
// 	XsiType   string `xml:"xsi:type,attr"`
// 	// UseAdGroup bool `xml:"useAdGroup"`
// 	OptIn bool `xml:"optIn"`
// }

type AdgroupOperations struct {
	XMLName                 xml.Name `xml:"operations"`
	Operator                string `xml:"operator"`
	Operand AdgroupOperand `xml:"operand"`
	// Name                    string `xml:"operand>name"`
	// Status                  string `xml:"operand>status"`
	// Period                  string `xml:"operand>budget>period"`
	// MicroAmount             string `xml:"operand>budget>amount>microAmount"`
	// DeliveryMethod          string `xml:"operand>budget>deliveryMethod"`
	// BiddingStrategy         BiddingStrategy `xml:"operand>biddingStrategy"`
	// Settings Settings       `xml:"operand>settings"`
	// TargetGoogleSearch      bool `xml:"operand>networkSetting>targetGoogleSearch"`
	// TargetSearchNetwork     bool `xml:"operand>networkSetting>targetSearchNetwork"`
	// TargetContentNetwork    bool `xml:"operand>networkSetting>targetContentNetwork"`
	// TargetContentContextual bool `xml:"operand>networkSetting>targetContentContextual"`
}

type AdgroupOperand struct {
	Id           int64  `xml:"id"`
	CampaignId   int64  `xml:"campaignId"`
	CampaignName string `xml:"campaignName"`
	Name         string `xml:"name"`
	Status       string `xml:"status"`
}

func (self *adgroupService) Mutate(v AdgroupOperations) {
	// v.BiddingStrategy.Cm = "https://adwords.google.com/api/adwords/cm/" + self.conn.Version
	// v.BiddingStrategy.Xsi = "http://www.w3.org/2001/XMLSchema-instance"
	
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
	
	println("https://adwords" + self.conn.sandboxUrl + ".google.com/api/adwords/cm/" + self.conn.Version + "/AdGroupService")
	req, err := http.NewRequest("POST", 
		"https://adwords" + self.conn.sandboxUrl + ".google.com/api/adwords/cm/" + self.conn.Version + "/AdGroupService", 
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

