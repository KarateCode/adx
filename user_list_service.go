package adx

import (
	// "fmt"
	"errors"
	"encoding/xml"
	// "io"
	// "os"
)

type userListService struct {
	conn *Conn
}

type UserListSelector struct {
	XMLName   xml.Name `xml:"serviceSelector"`
	Fields []string `xml:"fields"`
	Predicates []Predicate `xml:"predicates"`
	// StartIndex int `xml:"paging>startIndex"`
	// NumberResults int `xml:"paging>numberResults"`
}

type UserListGet struct {
	XMLName   xml.Name `xml:"Envelope"`
	Body struct {
		Fault Fault
		XMLName   xml.Name
		GetResponse struct {
			// XMLName   xml.Name `xml:"getResponse"`
			Rval struct {
				// XMLName   xml.Name
				TotalNumEntries int `xml:"totalNumEntries"`
				Entries []struct {
					Id int64 `xml:"id"`
					Name string `xml:"name"`
					// Status string `xml:"status"`
					// Stats struct {
					// 	Network string `xml:"network"`
					// 	StatsType string `xml:"Stats.Type"`
					// } `xml:"stats"`
				} `xml:"entries"`
			} `xml:"rval"`
		} `xml:"getResponse"`
	}
}

func (self *userListService) Get(v UserListSelector) (*UserListGet, error) {
	dataGet := new(UserListGet)
	
	returnBody, err := CallApi(v, self.conn, "UserListService", "get")
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
