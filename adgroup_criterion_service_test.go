package adx

import (
	// "fmt"
	"testing"
	. "github.com/KarateCode/helpers"
)

func TestGetAdgroupCriterion(*testing.T) {
	adwords := New(AdxPull)
	
	data := AdgroupCriterionSelector{
		Fields:   []string{"Id", "Status", "MaxCpm", "AdGroupName"}, 
		
		Predicates: []Predicate{
			Predicate{
				Field:    "AdGroupId", 
				Operator: "IN", 
				Values:   []string{"2765812624"},
			},
			Predicate{
				Field:    "Id", // It's a vertical ID
				Operator: "IN", 
				Values:   []string{"5012832132"}, 
			},
		},
		
		StartIndex: 0, 
		NumberResults: 5000,
	}
	
	adcGet, err := adwords.AdgroupCriterionService.Get(data)
	if err != nil {
		panic(err)
	}
	
	// fmt.Printf("\nadcGet%+v\n", adcGet.Body.GetResponse.Rval.Entries[0])
	// println(adcGet.Body.GetResponse.Rval.Entries)
	ShouldEqual(1, adcGet.Body.GetResponse.Rval.TotalNumEntries)
}

// func TestAddAdgroupCriterion(*testing.T) {
// 	adwords := New(AdxPush)
// 	
// 	data := AdgroupCriterionMutateOperations{
// 		// Name:                    "Campaign Test 2", 
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
// 	adwords.AdgroupCriterionService.Mutate(data)
// }

