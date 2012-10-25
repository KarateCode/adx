package v201109

import (
	// "text/template"
	"encoding/xml"
	"errors"
	// "io"
	// "os"
	"github.com/KarateCode/adx"
)

var version = "v201109"

type Adwords struct {
	CampaignService          campaignService
	ServicedAccountService   servicedAccountService
	AdgroupCriterionService  adgroupCriterionService
	AdgroupService           adgroupService
	ConversionTrackerService conversionTrackerService
	UserListService          userListService
	ConstantDataService      constantDataService
	BulkMutateJobService     bulkMutateJobService
}

func New(auth adx.Auth) *Adwords {
	auth.Version = version
	conn := adx.Conn{Auth:auth, Token:adx.Authenticate(auth.Email, auth.Password)}
	adwords := Adwords{
		CampaignService:          campaignService(conn),
		ServicedAccountService:   servicedAccountService(conn),
		AdgroupCriterionService:  adgroupCriterionService(conn),
		AdgroupService:           adgroupService(conn),
		ConversionTrackerService: conversionTrackerService(conn),
		UserListService:          userListService(conn),
		ConstantDataService:      constantDataService(conn),
		BulkMutateJobService:     bulkMutateJobService(conn),
	}
	
	if auth.Sandbox {
		conn.SandboxUrl = "-sandbox"
	}
	
	return &adwords
}

type campaignService adx.Conn
// type campaignService struct {
// 	conn *Conn
// }

type CampaignGetSelector struct {
	XMLName   xml.Name `xml:"serviceSelector"`
	// XsiType   string `xml:"xsi:type,attr"`
	Fields []string `xml:"fields"`
	Field string `xml:"predicates>field"`
	Operator string `xml:"predicates>operator"`
	Values []string `xml:"predicates>values"`
}

type CampaignGet struct {
	XMLName   xml.Name `xml:"Envelope"`
	Fault struct {
		FaultCode string `xml:"faultcode"`
		FaultString string `xml:"faultstring"`
	}
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

func (self *campaignService) Get(v CampaignGetSelector) (*CampaignGet, error) {
	campaignGet := new(CampaignGet)
	
	returnBody, err := adx.CallApi(v, (*adx.Conn)(self), "CampaignService", "get")
	if err != nil {return nil, err}
	defer returnBody.Close()
	
	decoder := xml.NewDecoder(returnBody)
	err = decoder.Decode(campaignGet)
	if err != nil {return nil, err}
	
	if campaignGet.Fault.FaultString != "" {
		return nil, errors.New(campaignGet.Fault.FaultString)
	}
	
	// io.Copy(os.Stdout, returnBody) // uncomment this to view http response. Found a 414 once
	return campaignGet, nil
}

type BiddingStrategy struct {
	// XMLName   xml.Name `xml:"operations"`
	XsiType   string `xml:"xsi:type,attr"`
	Cm string `xml:"xmlns:cm,attr"` 
	Xsi string `xml:"xmlns:xsi,attr"`
}

type Settings struct {
	XsiType   string `xml:"xsi:type,attr"`
	// UseAdGroup bool `xml:"useAdGroup"`
	OptIn bool `xml:"optIn"`
}

type CampaignMutateOperations struct {
	XMLName                 xml.Name `xml:"operations"`
	Operator                string `xml:"operator"`
	Name                    string `xml:"operand>name"`
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

func (self *campaignService) Mutate(v CampaignMutateOperations) error {
	// v.BiddingStrategy.Cm = "https://adwords.google.com/api/adwords/cm/" + self.conn.Version
	// v.BiddingStrategy.Xsi = "http://www.w3.org/2001/XMLSchema-instance"
	// v := servicedAccountServiceGet{EnablePaging:false, SubmanagersOnly:false}

	returnBody, err := adx.CallApi(v, (*adx.Conn)(self), "CampaignService", "mutate")
	if err != nil {return err}
	defer returnBody.Close()
	// io.Copy(os.Stdout, res.Body) // uncomment this to view http response. Found a 414 once
	return nil
}

