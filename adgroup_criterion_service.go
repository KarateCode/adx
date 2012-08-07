package adx

import (
	// "fmt"
	// "text/template"
	"encoding/xml"
	// "io"
	// "os"
)

type adgroupCriterionService struct {
	conn *Conn
}

type Predicate struct {
	Field string `xml:"field"`
	Operator string `xml:"operator"`
	Values []string `xml:"values"`
}

type AdgroupCriterionSelector struct {
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

type AdgroupCriterionGet struct {
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

func (self *adgroupCriterionService) Get(v AdgroupCriterionSelector) (*AdgroupCriterionGet, error) {
	adgroupGet := new(AdgroupCriterionGet)
	
	returnBody, err := CallApi(v, self.conn, "AdGroupCriterionService", "get")
	if err != nil {return nil, err}
	defer returnBody.Close()
	
	decoder := xml.NewDecoder(returnBody)
	err = decoder.Decode(adgroupGet)
	if err != nil {return nil, err}
	
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

type AdgroupCriterionOperations struct {
	XMLName                 xml.Name `xml:"operations"`
	Operator                string `xml:"operator"`
	// Name                    string `xml:"operand>name"`
	// Status                  string `xml:"operand>status"`
	Period                  string `xml:"operand>budget>period"`
	MicroAmount             string `xml:"operand>budget>amount>microAmount"`
	DeliveryMethod          string `xml:"operand>budget>deliveryMethod"`
	BiddingStrategy         BiddingStrategy `xml:"operand>biddingStrategy"`
	// Settings Settings       `xml:"operand>settings"`
	// TargetGoogleSearch      bool `xml:"operand>networkSetting>targetGoogleSearch"`
	// TargetSearchNetwork     bool `xml:"operand>networkSetting>targetSearchNetwork"`
	// TargetContentNetwork    bool `xml:"operand>networkSetting>targetContentNetwork"`
	// TargetContentContextual bool `xml:"operand>networkSetting>targetContentContextual"`
}

func (self *adgroupCriterionService) Mutate(v AdgroupCriterionOperations) error {
	// v.BiddingStrategy.Cm = "https://adwords.google.com/api/adwords/cm/" + self.conn.Version
	// v.BiddingStrategy.Xsi = "http://www.w3.org/2001/XMLSchema-instance"
	// v := servicedAccountServiceGet{EnablePaging:false, SubmanagersOnly:false}
	
	returnBody, err := CallApi(v, self.conn, "CampaignService", "mutate")
	if err != nil {return err}
	defer returnBody.Close()
	
	// io.Copy(os.Stdout, res.Body) // uncomment this to view http response. Found a 414 once
	return nil
}

