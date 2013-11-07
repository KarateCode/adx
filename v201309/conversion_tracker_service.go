package v201309

import (
	// "fmt"
	"encoding/xml"
	"errors"
	// "io"
	// "os"
)

type conversionTrackerService Conn

type ConversionTrackerSelector struct {
	XMLName    xml.Name    `xml:"serviceSelector"`
	Fields     []string    `xml:"fields"`
	Predicates []Predicate `xml:"predicates"`
	// StartIndex int `xml:"paging>startIndex"`
	// NumberResults int `xml:"paging>numberResults"`
}

type ConversionTrackerEntry struct {
	Id     int64  `xml:"id"`
	Name   string `xml:"name"`
	Status string `xml:"status"`
	Stats  struct {
		Network   string `xml:"network"`
		StatsType string `xml:"Stats.Type"`
	} `xml:"stats"`
}

type ConversionTrackerGet struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		Fault       Fault
		XMLName     xml.Name
		GetResponse struct {
			// XMLName   xml.Name `xml:"getResponse"`
			Rval struct {
				// XMLName   xml.Name
				TotalNumEntries int                      `xml:"totalNumEntries"`
				Entries         []ConversionTrackerEntry `xml:"entries"`
			} `xml:"rval"`
		} `xml:"getResponse"`
	}
}

func (self *conversionTrackerService) Get(v ConversionTrackerSelector) (*ConversionTrackerGet, error) {
	dataGet := new(ConversionTrackerGet)

	returnBody, err := CallApi(v, (*Conn)(self), "ConversionTrackerService", "get", "cm")
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
