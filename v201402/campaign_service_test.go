package v201402

import (
	// "fmt"
	. "github.com/KarateCode/helpers"
	"testing"
)

func TestGetCampaign(*testing.T) {
	adwords := New(AdxPush)

	data := CampaignGetSelector{
		Fields:   []string{"Id", "Name"},
		Field:    "Id",
		Operator: "EQUALS",
		Values:   []string{"702011"},
	}
	campaignGet, err := adwords.CampaignService.Get(data)
	if err != nil {
		panic(err)
	}

	ShouldEqual(1, len(campaignGet.Body.GetResponse.Rval.Entries))
	ShouldEqual("Hello World! with cURL", campaignGet.Body.GetResponse.Rval.Entries[0].Name)
}

// func TestAddCampaign(*testing.T) {
// 	adwords := New(AdxPush)
// 	
// 	data := CampaignMutateOperations{
// 		Name:                    "Campaign Test 2", 
// 		// Status:                  "PAUSED", 
// 		Operator:                "ADD", 
// 		Period:                  "DAILY",
// 		MicroAmount:             "1000000000",
// 		DeliveryMethod:          "STANDARD",
// 		BiddingStrategy:         BiddingStrategy{XsiType:"cm:ManualCPC"},
// 		// Settings:                Settings{XsiType:"RealTimeBiddingSetting", OptIn:true},
// 		// Settings:                Settings{XsiType:"KeywordMatchSetting", OptIn:true},
// 		// Settings:                Settings{XsiType:"TargetRestrictSetting"},
// 		// TargetGoogleSearch:      false,
// 		// TargetSearchNetwork:     false,
// 		// TargetContentNetwork:    true,
// 		// TargetContentContextual: false,
// 	}
// 	adwords.CampaignService.Mutate(data)
// }
