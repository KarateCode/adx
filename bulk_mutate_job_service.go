package adx

import (
	// "fmt"
	// "io"; "os"
	"errors"
	"encoding/xml"
)

type bulkMutateJobService struct {
	conn *Conn
}

type BulkMutateJobSelector struct {
	XMLName   xml.Name `xml:"selector"`
	XsiType  string    `xml:"xsi:type,attr"`
	JobIds   []int64   `xml:"jobIds"`
	// Fields []string `xml:"fields"`
	// Predicates []Predicate `xml:"predicates"`
	// Ordering
	// StartIndex int `xml:"paging>startIndex"`
	// NumberResults int `xml:"paging>numberResults"`
}

type BulkMutateJobGet struct {
	XMLName   xml.Name `xml:"Envelope"`
	Body struct {
		Fault Fault
		XMLName   xml.Name
		GetResponse struct {
			// XMLName   xml.Name `xml:"getResponse"`
			Rval struct {
				// XMLName   xml.Name
				Id           int64  `xml:"id"`
				Status       string `xml:"status"`
				// TotalNumEntries int `xml:"totalNumEntries"`
				// Entries []struct {
				// 	Id int64 `xml:"id,omitempty"`
				// 	Name string `xml:"name"`
				// 	CampaignId int64 `xml:"campaignId"`
				// 	CampaignName string `xml:"campaignName"`
				// 	Status string `xml:"status"`
				// 	Bids struct {
				// 		AdGroupBidsType string `xml:"AdGroupBids.Type"`	
				// 		EnhancedCpcEnabled bool `xml:"enhancedCpcEnabled"`
				// 	} `xml:"bids"`
				// 	Stats struct {
				// 		Network string `xml:"network"`
				// 		StatsType string `xml:"Stats.Type"`
				// 	} `xml:"stats"`
				// } `xml:"entries"`
			} `xml:"rval"`
		} `xml:"getResponse"`
	}
}

func (self *bulkMutateJobService) Get(v BulkMutateJobSelector) (*BulkMutateJobGet, error) {
	adgroupGet := new(BulkMutateJobGet)
	
	returnBody, err := CallApi(v, self.conn, "MutateJobService", "get")
	if err != nil {return nil, err}
	defer returnBody.Close()
	
	// io.Copy(os.Stdout, returnBody)
	
	// ba, err := ioutil.ReadAll(returnBody)
	// if err != nil { panic(err)}
	// println(string(ba))
	// decoder := xml.NewDecoder(bytes.NewBufferString(string(ba)))
	
	decoder := xml.NewDecoder(returnBody)
	err = decoder.Decode(adgroupGet)
	if err != nil {return nil, err}
	
	if adgroupGet.Body.Fault.FaultString != "" {
		return nil, errors.New(adgroupGet.Body.Fault.FaultString)
	}
	// fmt.Printf("\nadgroupGet from AdgroupService %+v\n", adgroupGet)
	return adgroupGet, nil
}


type BulkMutateJobOperations struct {
	XMLName  xml.Name             `xml:"operations"`
	XsiType  string               `xml:"xsi:type,attr"`
	Operator string               `xml:"operator"`
	Operand  BulkMutateJobOperand `xml:"operand"`
}

type BulkMutateJobOperand struct {
	XsiType      string `xml:"xsi:type,attr"`
	Id           int64  `xml:"id,omitempty"`
	CampaignId   int64  `xml:"campaignId,omitempty"`
	CampaignName string `xml:"campaignName,omitempty"`
	Name         string `xml:"name,omitempty"`
	Status       string `xml:"status,omitempty"`
	
	AdgroupId int64 `xml:"adGroupId"`
	CriterionUse string `xml:"criterionUse,omitempty"`
	Criterion Criterion  `xml:"criterion"`
	// Criterion struct {
	// 	Id int64 `xml:"id"`
	// 	Url string `xml:"url"`
	// 	Type string `xml:"type"`
	// } `xml:"criterion"`
	UserStatus string `xml:"userStatus"`
}

type BulkMutateJobResponse struct {
	XMLName   xml.Name `xml:"Envelope"`
	Body struct {
		Fault Fault
		MutateResponse struct {
			Rval struct {
				Id     int64  `xml:"id"`
				Status string `xml:"status"`
			} `xml:"rval"`
		} `xml:"mutateResponse"`
	}
}

func (self *bulkMutateJobService) Mutate(v []*BulkMutateJobOperations) (*BulkMutateJobResponse, error) {
	mutate := new(BulkMutateJobResponse)
	
	returnBody, err := CallApi(v, self.conn, "MutateJobService", "mutate")
	if err != nil {return nil, err}
	defer returnBody.Close()
	
	// io.Copy(os.Stdout, returnBody) // uncomment this to view http response. Found a 414 once
	
	decoder := xml.NewDecoder(returnBody)
	err = decoder.Decode(mutate)
	if err != nil {return nil, err}
	// fmt.Printf("\nadgroupMutate%+v\n", mutate)
	
	if mutate.Body.Fault.FaultString != "" {
		return nil, errors.New(mutate.Body.Fault.FaultString)
	}
	
	return mutate, nil
}
