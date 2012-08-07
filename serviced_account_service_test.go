package adx

import (
	// "fmt"
	"testing"
	. "github.com/KarateCode/helpers"
)

func TestServicedAccountGet(*testing.T) {
	adxPush := AdxPush
	adxPush.Version = "v201109"
	
	adwords := New(adxPush)
	// fmt.Printf("\nAdxPush%+v\n", AdxPush)
	data := ServicedAccountGetSelector{EnablePaging:false, SubmanagersOnly:false}
	accounts, links, err := adwords.ServicedAccountService.Get(data)
	if err != nil {
		panic(err) 
	}
	ShouldEqual(5, len(*accounts))
	ShouldEqual(5, len(*links))
	
	adwords = New(AdxPull)
	data = ServicedAccountGetSelector{EnablePaging:false, SubmanagersOnly:false}
	accounts, links, err = adwords.ServicedAccountService.Get(data)
	if err != nil {
		panic(err) 
	}
	ShouldEqual(1, len(*accounts))
	ShouldEqual(0, len(*links))
}