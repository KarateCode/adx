package v201402

import (
	// "fmt"
	"errors"
	// "io"
	// "os"
	// "bytes"
	// "io/ioutil"
	// "text/template"
	"encoding/xml"
)

type adgroupService Conn

// type AdgroupPredicate struct {
// 	Field string `xml:"field"`
// 	Operator string `xml:"operator"`
// 	Values []string `xml:"values"`
// }

type AdgroupSelector struct {
	XMLName xml.Name `xml:"serviceSelector"`
	// XsiType   string `xml:"xsi:type,attr"`
	Fields []string `xml:"fields"`
	// Field string `xml:"predicates>field"`
	// Operator string `xml:"predicates>operator"`
	// Values []string `xml:"predicates>values"`
	Predicates []Predicate `xml:"predicates"`
	// Ordering Ordering `xml:"ordering,omitempty"`
	Ordering
	StartIndex    int `xml:"paging>startIndex"`
	NumberResults int `xml:"paging>numberResults"`
}

type AdgroupEntry struct {
	Id           int64  `xml:"id"`
	Name         string `xml:"name"`
	CampaignId   int64  `xml:"campaignId"`
	CampaignName string `xml:"campaignName"`
	Status       string `xml:"status"`
	Bids         struct {
		AdGroupBidsType    string `xml:"AdGroupBids.Type"`
		EnhancedCpcEnabled bool   `xml:"enhancedCpcEnabled"`
	} `xml:"bids"`
	Stats struct {
		Network   string `xml:"network"`
		StatsType string `xml:"Stats.Type"`
	} `xml:"stats"`
}

type AdgroupGet struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		Fault       Fault
		XMLName     xml.Name
		GetResponse struct {
			// XMLName   xml.Name `xml:"getResponse"`
			Rval struct {
				// XMLName   xml.Name
				TotalNumEntries int            `xml:"totalNumEntries"`
				Entries         []AdgroupEntry `xml:"entries"`
			} `xml:"rval"`
		} `xml:"getResponse"`
	}
}

func (self *adgroupService) Get(v AdgroupSelector) (*AdgroupGet, error) {
	adgroupGet := new(AdgroupGet)

	returnBody, err := CallApi(v, (*Conn)(self), "AdGroupService", "get", "cm")
	if err != nil {
		return nil, err
	}
	defer returnBody.Close()

	// io.Copy(os.Stdout, returnBody)

	// ba, err := ioutil.ReadAll(returnBody)
	// if err != nil { panic(err)}
	// println(string(ba))
	// decoder := xml.NewDecoder(bytes.NewBufferString(string(ba)))

	decoder := xml.NewDecoder(returnBody)
	err = decoder.Decode(adgroupGet)
	if err != nil {
		return nil, err
	}

	if adgroupGet.Body.Fault.FaultString != "" {
		return nil, errors.New(adgroupGet.Body.Fault.FaultString)
	}
	// fmt.Printf("\nadgroupGet from AdgroupService %+v\n", adgroupGet)
	return adgroupGet, nil
}

type AdgroupOperations struct {
	XMLName  xml.Name       `xml:"operations"`
	Operator string         `xml:"operator"`
	Operand  AdgroupOperand `xml:"operand"`
}

type AdgroupOperand struct {
	Id           int64  `xml:"id"`
	CampaignId   int64  `xml:"campaignId"`
	CampaignName string `xml:"campaignName"`
	Name         string `xml:"name,omitempty"`
	Status       string `xml:"status"`
}

func (self *adgroupService) Mutate(v AdgroupOperations) error {
	// v.BiddingStrategy.Cm = "https://adwords.google.com/api/adwords/cm/" + self.conn.Version
	// v.BiddingStrategy.Xsi = "http://www.w3.org/2001/XMLSchema-instance"	
	// v := servicedAccountServiceGet{EnablePaging:false, SubmanagersOnly:false}
	adgroupMutate := new(MutateResponse)

	returnBody, err := CallApi(v, (*Conn)(self), "AdGroupService", "mutate", "cm")
	if err != nil {
		return err
	}
	defer returnBody.Close()

	decoder := xml.NewDecoder(returnBody)
	err = decoder.Decode(adgroupMutate)
	if err != nil {
		return err
	}
	// fmt.Printf("\nadgroupMutate%+v\n", adgroupMutate)

	if adgroupMutate.Body.Fault.FaultString != "" {
		return errors.New(adgroupMutate.Body.Fault.FaultString)
	}

	// io.Copy(os.Stdout, returnBody) // uncomment this to view http response. Found a 414 once
	return nil
}
