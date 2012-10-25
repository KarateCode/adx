package v201109

import (
	// "fmt"
	"testing"
	. "github.com/KarateCode/helpers"
	"github.com/KarateCode/adx"
)

func TestServicedAccountGet(*testing.T) {
	adxPush := adx.AdxPush
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
	
	adwords = New(adx.AdxPull)
	data = ServicedAccountGetSelector{EnablePaging:false, SubmanagersOnly:false}
	accounts, links, err = adwords.ServicedAccountService.Get(data)
	if err != nil {
		panic(err) 
	}
	ShouldEqual(1, len(*accounts))
	ShouldEqual(0, len(*links))
}