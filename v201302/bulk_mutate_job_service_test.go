package v201302

import (
	// "fmt"
	"testing"
	// "time"
	. "github.com/KarateCode/helpers"
)

func TestGetBulkMutateJob(*testing.T) {
	// adwords := New(AdxPull)
	adwords := New(AdxPush)

	data := BulkMutateJobSelector{
		XsiType: "BulkMutateJobSelector",
		JobIds:  []int64{-4857149772375705549},
	}

	adcGet, err := adwords.BulkMutateJobService.Get(data)
	if err != nil {
		panic(err)
	}

	// fmt.Printf("\nadcGet%+v\n", adcGet)
	ShouldEqual("COMPLETED", adcGet.Body.GetResponse.Rval.Status)
}

func TestGetResultBulkMutateJob(*testing.T) {
	adwords := New(AdxPush)

	data := BulkMutateJobSelector{
		XsiType: "BulkMutateJobSelector",
		JobIds:  []int64{8962009750808307540},
	}

	adcGet, err := adwords.BulkMutateJobService.GetResult(data)
	if err != nil {
		panic(err)
	}

	// fmt.Printf("\nadcGet%+v\n", adcGet)
	// for _, v := range adcGet.Body.GetResultResponse.Rval.SimpleMutateResult.Errors {
	// 	println("reason: ", v.Reason, "trigger: ", v.Trigger, "field path: ", v.FieldPath)
	// }

	ShouldEqual(20, len(adcGet.Body.GetResultResponse.Rval.SimpleMutateResult.Results))
	ShouldEqual(20, len(adcGet.Body.GetResultResponse.Rval.SimpleMutateResult.Errors))
}

func TestAddRemoveBulkMutateJob(*testing.T) {
	// AdxPush.Version = "v201109"
	adwords := New(AdxPush)
	// adgroupName := `Sample Adgroup ` + time.Now().String()
	var adcGet *BulkMutateJobGet
	// var err error

	// adcGet, err = adwords.BulkMutateJobService.Get(getData)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("\nadcGet%+v\n", adcGet.Body)
	// ShouldEqual(0, adcGet.Body.GetResponse.Rval.TotalNumEntries)
	var addData []*BulkMutateJobOperations
	// if adcGet.Body.GetResponse.Rval.TotalNumEntries == 0 {
	// println("Creating a new adgroup")
	// Create a new adgroup
	operation := BulkMutateJobOperations{
		XsiType:  "AdGroupCriterionOperation",
		Operator: "ADD",
		Operand: BulkMutateJobOperand{
			XsiType:   "BiddableAdGroupCriterion",
			AdgroupId: 2765812624,
			Criterion: Criterion{
				Id: 5012832132,
			},
			UserStatus: "ACTIVE",
			// Bids: Bids{
			// 	XsiType: "ManualCPMAdGroupCriterionBids",
			// 	MaxCpm: 320000,
			// },
		},
	}

	addData = append(addData, &operation)
	mutateResponse, err := adwords.BulkMutateJobService.Mutate(addData)
	if err != nil {
		panic(err)
	}
	// }

	getData := BulkMutateJobSelector{
		XsiType: "BulkMutateJobSelector",
		JobIds:  []int64{mutateResponse.Body.MutateResponse.Rval.Id},
		// Fields:   []string{"Id", "Status", "MaxCpm", "AdGroupName"}, 
		// Predicates: []Predicate{
		// 	Predicate{
		// 		Field:    "CampaignId", 
		// 		Operator: "EQUALS", 
		// 		Values:   []string{"702011"},
		// 	},
		// 	Predicate{
		// 		Field:    "Status", 
		// 		Operator: "NOT_IN", 
		// 		Values:   []string{"DELETED"},
		// 	},
		// },
		// StartIndex: 0, 
		// NumberResults: 5000,
	}

	// Make sure there is 1 active adgroups
	adcGet, err = adwords.BulkMutateJobService.Get(getData)
	if err != nil {
		panic(err)
	}
	// fmt.Printf("\nadcGet%+v\n", adcGet)
	// ShouldEqual(1, adcGet.Body.GetResponse.Rval.TotalNumEntries)
	ShouldEqual("COMPLETED", adcGet.Body.GetResponse.Rval.Status)

	// adgroupId := adcGet.Body.GetResponse.Rval.Entries[0].Id

	// Now remove the adgroup
	// removeData := BulkMutateJobOperations{
	// 	Operator:  "SET", 
	// 	Operand: BulkMutateJobOperand{
	// 		Id: adgroupId, 
	// 		CampaignId: 702011, 
	// 		Name: adgroupName,
	// 		Status: "DELETED",
	// 	},
	// }
	// err = adwords.BulkMutateJobService.Mutate(removeData)
	// if err != nil {
	// 	panic(err)
	// }
	// 
	// // Make sure there is 0 active adgroups
	// adcGet, err = adwords.BulkMutateJobService.Get(getData)
	// if err != nil {
	// 	panic(err)
	// }
	// // fmt.Printf("\nadcGet%+v\n", adcGet.Body)
	// ShouldEqual(0, adcGet.Body.GetResponse.Rval.TotalNumEntries)
}
