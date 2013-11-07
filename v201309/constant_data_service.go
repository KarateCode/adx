package v201309

import (
	// "fmt"
	"encoding/xml"
	"errors"
	// "io"
	// "os"
)

type constantDataService Conn

// type ConstantDataSelector struct {
// 	XMLName   xml.Name `xml:"serviceSelector"`
// 	Fields []string `xml:"fields"`
// 	Predicates []Predicate `xml:"predicates"`
// 	// StartIndex int `xml:"paging>startIndex"`
// 	// NumberResults int `xml:"paging>numberResults"`
// }

type ConstantDataGetVerticalCriterion struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		Fault                        Fault
		XMLName                      xml.Name
		GetVerticalCriterionResponse struct {
			// XMLName   xml.Name `xml:"getResponse"`
			Rval []struct {
				// XMLName   xml.Name
				CriterionType    string   `xml:"Criterion.Type"`
				VerticalId       int64    `xml:"verticalId"`
				VerticalParentId int64    `xml:"verticalParentId"`
				Paths            []string `xml:"path"`
			} `xml:"rval"`
		} `xml:"getVerticalCriterionResponse"`
	}
}

func (self *constantDataService) GetVerticalCriterion() (*ConstantDataGetVerticalCriterion, error) {
	dataGet := new(ConstantDataGetVerticalCriterion)

	returnBody, err := CallApi(nil, (*Conn)(self), "ConstantDataService", "getVerticalCriterion", "cm")
	if err != nil {
		return nil, err
	}
	defer returnBody.Close()

	// io.Copy(os.Stdout, returnBody)

	decoder := xml.NewDecoder(returnBody)
	err = decoder.Decode(dataGet)
	if err != nil {
		return nil, err
	}

	if dataGet.Body.Fault.FaultString != "" {
		return nil, errors.New(dataGet.Body.Fault.FaultString)
	}
	// fmt.Printf("\nadgroupGet from AdgroupService %+v\n", dataGet)
	return dataGet, nil
}
