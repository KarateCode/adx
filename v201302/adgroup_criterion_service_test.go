package v201302

import (
	// "fmt"
	. "github.com/KarateCode/helpers"
	// "strconv"
	"testing"
	// "time"
)

func TestGetAdgroupCriterion(*testing.T) {
	adwords := New(AdxPull)

	data := AdgroupCriterionSelector{
		Fields: []string{"Id", "Status", "MaxCpm", "AdGroupName"},

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

		StartIndex:    0,
		NumberResults: 5000,
	}

	adcGet, err := adwords.AdgroupCriterionService.Get(data)
	if err != nil {
		panic(err)
	}

	// fmt.Printf("\nadcGet%+v\n", adcGet.Body.GetResponse.Rval.Entries[0])
	ShouldEqual(1, adcGet.Body.GetResponse.Rval.TotalNumEntries)
}

func TestSetMaxCpm(*testing.T) {
	adwords := New(AdxPull)

	// Update placement
	createData := AdgroupCriterionOperations{
		Operator: "SET",
		Operand: AdgroupCriterionOperand{
			XsiType:   "BiddableAdGroupCriterion",
			AdgroupId: 2765812624,
			// CriterionUse: "BIDDABLE",
			Criterion: Criterion{
				// Type: "VERTICAL", 
				Id:      5012832132,
				XsiType: "Placement",
				// Url: "http://pluto.google.com",
			},
			UserStatus: "ACTIVE",
			Bids: Bids{
				XsiType: "ManualCPMAdGroupCriterionBids",
				// BidsXsiType: "ManualCPMAdGroupCriterionBids",
				MaxCpm: 320000,
				// MaxCpm: MaxCpm{Amount: Amount{MicroAmount:320000}},
			},
		},
	}

	// for i := 0; i<20; i++ {  // Why doesn't this cause a RATE EXCEEDED error?
	if err := adwords.AdgroupCriterionService.Mutate(createData); err != nil {
		panic(err)
	}
	// }

	// Read placement
	data := AdgroupCriterionSelector{
		Fields: []string{"Id", "Status", "MaxCpm", "AdGroupName"},

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

		StartIndex:    0,
		NumberResults: 5000,
	}
	adcGet, err := adwords.AdgroupCriterionService.Get(data)
	if err != nil {
		panic(err)
	}

	// fmt.Printf("\nadcGet%+v\n", adcGet.Body.GetResponse.Rval.Entries[0])
	ShouldEqual(1, adcGet.Body.GetResponse.Rval.TotalNumEntries)
	ShouldEqual("ACTIVE", adcGet.Body.GetResponse.Rval.Entries[0].UserStatus)
	// ShouldEqual(int64(320000), adcGet.Body.GetResponse.Rval.Entries[0].Bids.MaxCpm)
	ShouldEqual(int64(320000), adcGet.Body.GetResponse.Rval.Entries[0].Bids.MaxCpm.Amount.MicroAmount)
}

// func TestAddRemoveAdgroupCriterion(*testing.T) {
// 	adwords := New(AdxPush)
// 	adgroupName := `Sample Adgroup ` + time.Now().String()
// 	var adcGet *AdgroupGet
// 	var err error
// 	var adcsGet *AdgroupCriterionGet
// 	var adgroupId int64

// 	getData := AdgroupSelector{
// 		Fields: []string{"Id", "Status", "MaxCpm", "AdGroupName"},
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
// 		StartIndex:    0,
// 		NumberResults: 5000,
// 	}

// 	// Make sure there are 0 active adgroups
// 	adcGet, err = adwords.AdgroupService.Get(getData)
// 	if err != nil {
// 		panic(err)
// 	}
// 	// fmt.Printf("\nadcGet%+v\n", adcGet.Body)
// 	// ShouldEqual(0, adcGet.Body.GetResponse.Rval.TotalNumEntries)
// 	if adcGet.Body.GetResponse.Rval.TotalNumEntries == 0 {
// 		// Create a new adgroup
// 		addData := AdgroupOperations{
// 			Operator: "ADD",
// 			Operand: AdgroupOperand{
// 				CampaignId: 702011,
// 				Name:       adgroupName,
// 				Status:     "PAUSED",
// 			},
// 		}
// 		err = adwords.AdgroupService.Mutate(addData)
// 		if err != nil {
// 			panic(err)
// 		}
// 	}

// 	// Make sure there is 1 active adgroups
// 	adcGet, err = adwords.AdgroupService.Get(getData)
// 	if err != nil {
// 		panic(err)
// 	}
// 	// fmt.Printf("\nadcGet%+v\n", adcGet.Body)
// 	ShouldEqual(1, adcGet.Body.GetResponse.Rval.TotalNumEntries)
// 	adgroupId = adcGet.Body.GetResponse.Rval.Entries[0].Id
// 	// fmt.Printf("\nadgroupId%v\n", adgroupId)

// 	// Add first placement
// 	createData := AdgroupCriterionOperations{
// 		Operator: "ADD",
// 		Operand: AdgroupCriterionOperand{
// 			XsiType:   "BiddableAdGroupCriterion",
// 			AdgroupId: adgroupId,
// 			// CriterionUse: "BIDDABLE",
// 			Criterion: Criterion{
// 				XsiType: "Placement",
// 				// Keyword: "Comics",
// 				Url: "http://mars.google.com",
// 				// MatchType: "EXACT",
// 			},
// 			UserStatus: "ACTIVE",
// 			// Bids: Bids{
// 			// XsiType: "ManualCPMAdGroupCriterionBids",
// 			// MaxCpm: MaxCpm{Amount: Amount{MicroAmount:250000}},
// 			// },
// 		},
// 	}
// 	if err = adwords.AdgroupCriterionService.Mutate(createData); err != nil {
// 		panic(err)
// 	}

// 	// Add second placement
// 	createData = AdgroupCriterionOperations{
// 		Operator: "ADD",
// 		Operand: AdgroupCriterionOperand{
// 			XsiType:   "BiddableAdGroupCriterion",
// 			AdgroupId: adgroupId,
// 			// CriterionUse: "BIDDABLE",
// 			Criterion: Criterion{
// 				XsiType: "Placement",
// 				// Keyword: "Comics",
// 				Url: "http://pluto.google.com",
// 				// MatchType: "EXACT",
// 			},
// 			// UserStatus: "ACTIVE",
// 		},
// 	}
// 	if err = adwords.AdgroupCriterionService.Mutate(createData); err != nil {
// 		panic(err)
// 	}

// 	data := AdgroupCriterionSelector{
// 		Fields: []string{"Id", "Status", "MaxCpm", "AdGroupName", "PlacementUrl"},
// 		Predicates: []Predicate{
// 			Predicate{
// 				Field:    "AdGroupId",
// 				Operator: "IN",
// 				Values:   []string{strconv.FormatInt(adgroupId, 10)},
// 			},
// 			Predicate{
// 				Field:    "Status",
// 				Operator: "NOT_IN",
// 				Values:   []string{"DELETED"},
// 			},
// 		},
// 		StartIndex:    0,
// 		NumberResults: 5000,
// 	}

// 	// Make sure 2 placements are in the system
// 	adcsGet, err = adwords.AdgroupCriterionService.Get(data)
// 	if err != nil {
// 		panic(err)
// 	}
// 	// fmt.Printf("\n\nadcGet%+v\n", adcsGet)
// 	ShouldEqual(2, adcsGet.Body.GetResponse.Rval.TotalNumEntries)
// 	ShouldEqual(2, len(adcsGet.Body.GetResponse.Rval.Entries))

// 	// Remove first placement
// 	removePlacement1 := AdgroupCriterionOperations{
// 		Operator: "REMOVE",
// 		Operand: AdgroupCriterionOperand{
// 			XsiType:   "BiddableAdGroupCriterion",
// 			AdgroupId: adgroupId,
// 			Criterion: Criterion{
// 				Id: 11743038,
// 			},
// 		},
// 	}
// 	if err = adwords.AdgroupCriterionService.Mutate(removePlacement1); err != nil {
// 		panic(err)
// 	}

// 	// Should only be 1 at this point
// 	adcsGet, err = adwords.AdgroupCriterionService.Get(data)
// 	if err != nil {
// 		panic(err)
// 	}
// 	// fmt.Printf("\n\nadcGet%+v\n", adcsGet)
// 	ShouldEqual(1, adcsGet.Body.GetResponse.Rval.TotalNumEntries)
// 	ShouldEqual(1, len(adcsGet.Body.GetResponse.Rval.Entries))

// 	// Remove placement 2
// 	removePlacement2 := AdgroupCriterionOperations{
// 		Operator: "REMOVE",
// 		Operand: AdgroupCriterionOperand{
// 			XsiType:   "BiddableAdGroupCriterion",
// 			AdgroupId: adgroupId,
// 			Criterion: Criterion{
// 				Id: 30879662,
// 			},
// 		},
// 	}
// 	if err = adwords.AdgroupCriterionService.Mutate(removePlacement2); err != nil {
// 		panic(err)
// 	}

// 	// Should only 0 at this point
// 	adcsGet, err = adwords.AdgroupCriterionService.Get(data)
// 	if err != nil {
// 		panic(err)
// 	}
// 	// fmt.Printf("\n\nadcGet%+v\n", adcsGet)
// 	ShouldEqual(0, adcsGet.Body.GetResponse.Rval.TotalNumEntries)
// 	ShouldEqual(0, len(adcsGet.Body.GetResponse.Rval.Entries))

// 	// Now remove the adgroup
// 	removeData := AdgroupOperations{
// 		Operator: "SET",
// 		Operand: AdgroupOperand{
// 			Id:         adgroupId,
// 			CampaignId: 702011,
// 			Name:       adgroupName,
// 			Status:     "DELETED",
// 		},
// 	}
// 	err = adwords.AdgroupService.Mutate(removeData)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Make sure there is 0 active adgroups
// 	adcGet, err = adwords.AdgroupService.Get(getData)
// 	if err != nil {
// 		panic(err)
// 	}
// 	// fmt.Printf("\nadcGet%+v\n", adcGet.Body)
// 	ShouldEqual(0, adcGet.Body.GetResponse.Rval.TotalNumEntries)
// }

// func TestRetrieveAdgroupCriterion(*testing.T) {
// 	var adgroupId int64
// 	adgroupId = 3066700139
// 	adwords := New(AdxPush)
// 	var err error
// 	var adcsGet *AdgroupCriterionGet
// 	
// 	createData := AdgroupCriterionOperations{
// 		Operator:                "ADD", 
// 		Operand: AdgroupCriterionOperand{
// 			XsiType: "BiddableAdGroupCriterion",
// 			AdgroupId: adgroupId,
// 			// CriterionUse: "BIDDABLE",
// 			Criterion: Criterion{
// 				Type: "Placement", 
// 				// Keyword: "Comics",
// 				Url: "http://mars.google.com",
// 				// MatchType: "EXACT",
// 			},
// 			// UserStatus: "ACTIVE",
// 		},
// 	}
// 	if err = adwords.AdgroupCriterionService.Mutate(createData); err != nil {
// 		panic(err)
// 	}
// 	createData = AdgroupCriterionOperations{
// 		Operator:                "ADD", 
// 		Operand: AdgroupCriterionOperand{
// 			XsiType: "BiddableAdGroupCriterion",
// 			AdgroupId: adgroupId,
// 			// CriterionUse: "BIDDABLE",
// 			Criterion: Criterion{
// 				Type: "Placement", 
// 				// Keyword: "Comics",
// 				Url: "http://pluto.google.com",
// 				// MatchType: "EXACT",
// 			},
// 			// UserStatus: "ACTIVE",
// 		},
// 	}
// 	if err = adwords.AdgroupCriterionService.Mutate(createData); err != nil {
// 		panic(err)
// 	}
// 	
// 	
// 	
// 	data := AdgroupCriterionSelector{
// 		Fields:   []string{"Id", "Status", "MaxCpm", "AdGroupName", "PlacementUrl"}, 
// 		Predicates: []Predicate{
// 			Predicate{
// 				Field:    "AdGroupId", 
// 				Operator: "IN", 
// 				Values:   []string{string(adgroupId)},
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
// 	adcsGet, err = adwords.AdgroupCriterionService.Get(data)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("\n\nadcGet%+v\n", adcsGet)
// 	ShouldEqual(2, adcsGet.Body.GetResponse.Rval.TotalNumEntries)
// 	ShouldEqual(2, len(adcsGet.Body.GetResponse.Rval.Entries))
// 	
// 	
// 	// Remove placements
// 	removePlacement1 := AdgroupCriterionOperations{
// 		Operator:                "REMOVE", 
// 		Operand: AdgroupCriterionOperand{
// 			XsiType: "BiddableAdGroupCriterion",
// 			AdgroupId: adgroupId,
// 			Criterion: Criterion{
// 				Id: 11743038,
// 			},
// 		},
// 	}
// 	if err = adwords.AdgroupCriterionService.Mutate(removePlacement1); err != nil {
// 		panic(err)
// 	}
// 	
// 	// Should only be 1 at this point
// 	adcsGet, err = adwords.AdgroupCriterionService.Get(data)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("\n\nadcGet%+v\n", adcsGet)
// 	ShouldEqual(1, adcsGet.Body.GetResponse.Rval.TotalNumEntries)
// 	ShouldEqual(1, len(adcsGet.Body.GetResponse.Rval.Entries))
// 	
// 	// Remove placement 2
// 	removePlacement2 := AdgroupCriterionOperations{
// 		Operator:                "REMOVE", 
// 		Operand: AdgroupCriterionOperand{
// 			XsiType: "BiddableAdGroupCriterion",
// 			AdgroupId: adgroupId,
// 			Criterion: Criterion{
// 				Id: 30879662,
// 			},
// 		},
// 	}
// 	if err = adwords.AdgroupCriterionService.Mutate(removePlacement2); err != nil {
// 		panic(err)
// 	}
// 	
// 	// Should only 0 at this point
// 	adcsGet, err = adwords.AdgroupCriterionService.Get(data)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("\n\nadcGet%+v\n", adcsGet)
// 	ShouldEqual(0, adcsGet.Body.GetResponse.Rval.TotalNumEntries)
// 	ShouldEqual(0, len(adcsGet.Body.GetResponse.Rval.Entries))
// }
