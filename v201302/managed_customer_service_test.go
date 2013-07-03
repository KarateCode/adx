package v201302

import (
	// "fmt"
	. "github.com/KarateCode/helpers"
	"testing"
)

func TestManagedCustomerGet(*testing.T) {
	var adwords *Adwords
	var data ManagedCustomerGetSelector
	var accounts *[]Entry
	var links *[]Link
	var err error

	adxPush := AdxPush
	adxPush.Version = "v201109"
	adwords = New(adxPush)
	// fmt.Printf("\nAdxPush%+v\n", AdxPush)
	data = ManagedCustomerGetSelector{EnablePaging: false, SubmanagersOnly: false}
	accounts, links, err = adwords.ManagedCustomerService.Get(data)
	if err != nil {
		panic(err)
	}
	ShouldEqual(6, len(*accounts))
	ShouldEqual(5, len(*links))

	// adwords = New(AdxPull)
	// data = ManagedCustomerGetSelector{
	// 	// Fields: []string{"name", "login"},
	// 	Fields: []string{"Id", "Status", "MaxCpm", "AdGroupName"}, 
	// 	EnablePaging:false, 
	// 	SubmanagersOnly:false,
	// }
	// accounts, links, err = adwords.ManagedCustomerService.Get(data)
	// if err != nil {
	// 	panic(err) 
	// }
	// fmt.Printf("\naccounts%+v\n", accounts)
	// ShouldEqual(1, len(*accounts))
	// ShouldEqual(0, len(*links))
}
