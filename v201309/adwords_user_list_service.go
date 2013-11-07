package v201309

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"os"
)

type adwordsUserListService Conn

type AdwordsUserListSelector struct {
	XMLName    xml.Name    `xml:"serviceSelector"`
	Fields     []string    `xml:"fields"`
	Predicates []Predicate `xml:"predicates"`
	// StartIndex int `xml:"paging>startIndex"`
	// NumberResults int `xml:"paging>numberResults"`
}

type AdwordsUserListEntry struct {
	Id   int64  `xml:"id"`
	Name string `xml:"name"`
}

type AdwordsUserListGet struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		Fault       Fault
		XMLName     xml.Name
		GetResponse struct {
			// XMLName   xml.Name `xml:"getResponse"`
			Rval struct {
				// XMLName   xml.Name
				TotalNumEntries int                    `xml:"totalNumEntries"`
				Entries         []AdwordsUserListEntry `xml:"entries"`
			} `xml:"rval"`
		} `xml:"getResponse"`
	}
}

func (self *adwordsUserListService) Get(v AdwordsUserListSelector) (*AdwordsUserListGet, error) {
	fmt.Printf("v: %+v\n", v)
	dataGet := new(AdwordsUserListGet)

	returnBody, err := CallApi(v, (*Conn)(self), "AdwordsUserListService", "get", "rm")
	if err != nil {
		return nil, err
	}
	defer returnBody.Close()

	io.Copy(os.Stdout, returnBody)
	println("*")
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
