package v201206

import (
	// "fmt"
	"errors"
	"encoding/xml"
	// "io"
	// "os"
	"github.com/KarateCode/adx"
)

type userListService adx.Conn

type UserListSelector struct {
	XMLName   xml.Name `xml:"serviceSelector"`
	Fields []string `xml:"fields"`
	Predicates []adx.Predicate `xml:"predicates"`
	// StartIndex int `xml:"paging>startIndex"`
	// NumberResults int `xml:"paging>numberResults"`
}

type UserListEntry struct {
	Id int64 `xml:"id"`
	Name string `xml:"name"`
}

type UserListGet struct {
	XMLName   xml.Name `xml:"Envelope"`
	Body struct {
		Fault adx.Fault
		XMLName   xml.Name
		GetResponse struct {
			// XMLName   xml.Name `xml:"getResponse"`
			Rval struct {
				// XMLName   xml.Name
				TotalNumEntries int `xml:"totalNumEntries"`
				Entries []UserListEntry `xml:"entries"`
			} `xml:"rval"`
		} `xml:"getResponse"`
	}
}

func (self *userListService) Get(v UserListSelector) (*UserListGet, error) {
	dataGet := new(UserListGet)
	
	returnBody, err := adx.CallApi(v, (*adx.Conn)(self), "UserListService", "get")
	if err != nil {return nil, err}
	defer returnBody.Close()
	
	// io.Copy(os.Stdout, returnBody)
	
	decoder := xml.NewDecoder(returnBody)
	err = decoder.Decode(dataGet)
	if err != nil {return nil, err}
	
	if dataGet.Body.Fault.FaultString != "" {
		return nil, errors.New(dataGet.Body.Fault.FaultString)
	}
	// fmt.Printf("\nadgroupGet from AdgroupService %+v\n", dataGet)
	return dataGet, nil
}
