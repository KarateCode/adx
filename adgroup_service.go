package adx

import (
	// "fmt"
	// "text/template"
	"encoding/xml"
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
	
	returnBody, err := CallApi(v, self.conn, "AdGroupService", "get")
	if err != nil {return nil, err}
	defer returnBody.Close()
	
	decoder := xml.NewDecoder(returnBody)
	err = decoder.Decode(adgroupGet)
	if err != nil {return nil, err}
	
	// fmt.Printf("\nadgroupGet%v\n", adgroupGet)
	// io.Copy(os.Stdout, res.Body) // uncomment this to view http response. Found a 414 once
	return adgroupGet, nil
}


type AdgroupOperations struct {
	XMLName                 xml.Name `xml:"operations"`
	Operator                string `xml:"operator"`
	Operand AdgroupOperand `xml:"operand"`
}

type AdgroupOperand struct {
	Id           int64  `xml:"id"`
	CampaignId   int64  `xml:"campaignId"`
	CampaignName string `xml:"campaignName"`
	Name         string `xml:"name"`
	Status       string `xml:"status"`
}

func (self *adgroupService) Mutate(v AdgroupOperations) error {
	// v.BiddingStrategy.Cm = "https://adwords.google.com/api/adwords/cm/" + self.conn.Version
	// v.BiddingStrategy.Xsi = "http://www.w3.org/2001/XMLSchema-instance"	
	// v := servicedAccountServiceGet{EnablePaging:false, SubmanagersOnly:false}
	
	returnBody, err := CallApi(v, self.conn, "AdGroupService", "mutate")
	if err != nil {return err}
	defer returnBody.Close()
	
	// io.Copy(os.Stdout, returnBody) // uncomment this to view http response. Found a 414 once
	return nil
}

