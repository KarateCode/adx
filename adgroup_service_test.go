package adx

import (
	// "fmt"
	"testing"
	"time"
	. "github.com/KarateCode/helpers"
)

func TestGetAdgroup(*testing.T) {
	adwords := New(AdxPull)
	
	data := AdgroupSelector{
		Fields:   []string{"Id", "Status", "MaxCpm", "AdGroupName"}, 
		
		Predicates: []Predicate{
			Predicate{
				Field:    "AdGroupId", 
				Operator: "IN", 
				Values:   []string{"2765812624"},
			},
		// 	Predicate{
		// 		Field:    "Id", // It's a vertical ID
		// 		Operator: "IN", 
		// 		Values:   []string{"5012832132"}, 
		// 	},
		},
		
		StartIndex: 0, 
		NumberResults: 5000,
	}
	
	adcGet, err := adwords.AdgroupService.Get(data)
	if err != nil {
		panic(err)
	}
	
	// fmt.Printf("\nadcGet%+v\n", adcGet.Body)
	// println(adcGet.Body.GetResponse.Rval.Entries)
	ShouldEqual(1, adcGet.Body.GetResponse.Rval.TotalNumEntries)
}

func TestAddRemoveAdgroup(*testing.T) {
	adwords := New(AdxPush)
	adgroupName := `Sample Adgroup ` + time.Now().String()
	var adcGet *AdgroupGet
	var err error
	
	getData := AdgroupSelector{
		Fields:   []string{"Id", "Status", "MaxCpm", "AdGroupName"}, 
		Predicates: []Predicate{
			Predicate{
				Field:    "CampaignId", 
				Operator: "EQUALS", 
				Values:   []string{"702011"},
			},
			Predicate{
				Field:    "Status", 
				Operator: "NOT_IN", 
				Values:   []string{"DELETED"},
			},
		},
		StartIndex: 0, 
		NumberResults: 5000,
	}
	
	// Make sure there are 0 active adgroups
	adcGet, err = adwords.AdgroupService.Get(getData)
	if err != nil {
		panic(err)
	}
	// fmt.Printf("\nadcGet%+v\n", adcGet.Body)
	ShouldEqual(0, adcGet.Body.GetResponse.Rval.TotalNumEntries)
	
	// Create a new adgroup
	addData := AdgroupOperations{
		Operator:  "ADD", 
		Operand: AdgroupOperand{
			CampaignId: 702011, 
			Name: adgroupName, 
			Status: "PAUSED",
		},
	}
	adwords.AdgroupService.Mutate(addData)
	
	// Make sure there is 1 active adgroups
	adcGet, err = adwords.AdgroupService.Get(getData)
	if err != nil {
		panic(err)
	}
	// fmt.Printf("\nadcGet%+v\n", adcGet.Body)
	ShouldEqual(1, adcGet.Body.GetResponse.Rval.TotalNumEntries)
	
	adgroupId := adcGet.Body.GetResponse.Rval.Entries[0].Id
	
	// Now remove the adgroup
	removeData := AdgroupOperations{
		Operator:  "SET", 
		Operand: AdgroupOperand{
			Id: adgroupId, 
			CampaignId: 702011, 
			Name: adgroupName,
			Status: "DELETED",
		},
	}
	adwords.AdgroupService.Mutate(removeData)
	
	// Make sure there is 0 active adgroups
	adcGet, err = adwords.AdgroupService.Get(getData)
	if err != nil {
		panic(err)
	}
	// fmt.Printf("\nadcGet%+v\n", adcGet.Body)
	ShouldEqual(0, adcGet.Body.GetResponse.Rval.TotalNumEntries)
}

 // func TestGetAdgroupByCampaign(*testing.T) {
 // 	// 3066697865
 // 	adwords := New(AdxPush)
 // 	data := AdgroupSelector{
 // 		Fields:     []string{"Id", "Status", "MaxCpm", "AdGroupName"}, 
 // 		Predicates: []Predicate{
 // 			Predicate{
 // 				Field:    "CampaignId", 
 // 				Operator: "EQUALS", 
 // 				Values:   []string{"702011"},
 // 			},
 // 			Predicate{
 // 				Field:    "Status", 
 // 				Operator: "NOT_IN", 
 // 				Values:   []string{"DELETED"},
 // 			},
 // 		},
 // 		StartIndex: 0, 
 // 		NumberResults: 5000,
 // 	}
 // 	
 // 	adcGet, err := adwords.AdgroupService.Get(data)
 // 	if err != nil {
 // 		panic(err)
 // 	}
 // 	fmt.Printf("\n\nadcGet%+v\n", adcGet)
 // 	ShouldEqual(1, adcGet.Body.GetResponse.Rval.TotalNumEntries)
 // 	
 // 	// removeData := AdgroupOperations{
 // 	// 	Operator:  "SET", 
 // 	// 	Operand: AdgroupOperand{
 // 	// 		Id: 3066697865, 
 // 	// 		CampaignId: 702011, 
 // 	// 		Name: "Sample Adgroup",
 // 	// 		Status: "DELETED",
 // 	// 	},
 // 	// }
 // 	// adwords.AdgroupService.Mutate(removeData)
 // 	// 
 // 	// 
 // 	// adcGet, err := adwords.AdgroupService.Get(data)
 // 	// if err != nil {
 // 	// 	panic(err)
 // 	// }
 // 	// 
 // 	// fmt.Printf("\n\nadcGet%+v\n", adcGet)
 // 	// // println(adcGet.Body.GetResponse.Rval.Entries)
 // 	// ShouldEqual(0, adcGet.Body.GetResponse.Rval.TotalNumEntries)
 // }

