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

func TestAddRemoveAdgroups(*testing.T) {
	// println("Testing Adgroup")
	
	AdxPush.Version = "v201109"
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
	// ShouldEqual(0, adcGet.Body.GetResponse.Rval.TotalNumEntries)
	
	if adcGet.Body.GetResponse.Rval.TotalNumEntries == 0 {
		// println("Creating a new adgroup")
		// Create a new adgroup
		addData := AdgroupOperations{
			Operator:  "ADD", 
			Operand: AdgroupOperand{
				CampaignId: 702011, 
				Name: adgroupName, 
				Status: "PAUSED",
			},
		}
		err = adwords.AdgroupService.Mutate(addData)
		if err != nil {
			panic(err)
		}
	}
	
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
	err = adwords.AdgroupService.Mutate(removeData)
	if err != nil {
		panic(err)
	}
	
	// Make sure there is 0 active adgroups
	adcGet, err = adwords.AdgroupService.Get(getData)
	if err != nil {
		panic(err)
	}
	// fmt.Printf("\nadcGet%+v\n", adcGet.Body)
	ShouldEqual(0, adcGet.Body.GetResponse.Rval.TotalNumEntries)
	// println("Testing Adgroup Complete")
}

// func TestRemoveAdgroupUtility(*testing.T) {
// 	adwords := New(AdxPush)
// 	
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
// 	adcGet, err := adwords.AdgroupService.Get(data)
// 	fmt.Printf("\n\nadcGet%+v\n", adcGet)
// 	if err != nil {
// 		panic(err)
// 	}
// 	
// 	if adcGet.Body.GetResponse.Rval.TotalNumEntries > 0 {
//  	removeData := AdgroupOperations{
//  		Operator:  "SET", 
//  		Operand: AdgroupOperand{
//  			Id: adcGet.Body.GetResponse.Rval.Entries[0].Id, 
//  			CampaignId: 702011, 
//  			Status: "DELETED",
//  		},
//  	}
//  	println("REMOVING")
//  	adwords.AdgroupService.Mutate(removeData)
// }
// 
// 	adcGet, err = adwords.AdgroupService.Get(data)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("\n\nadcGet%+v\n", adcGet)
// 	ShouldEqual(0, adcGet.Body.GetResponse.Rval.TotalNumEntries)
// }

