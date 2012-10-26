package v201206

import (
	"testing"
	"encoding/xml"
	"bytes"
	// "fmt"
	. "github.com/KarateCode/helpers"
)

func TestDecodingSoapFault(*testing.T) {
	adgroupGet := new(AdgroupGet)
	soapError := `<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
	<soap:Header>
		<ResponseHeader xmlns="https://adwords.google.com/api/adwords/cm/v201109">
			<requestId>0004c6d63c4669180aedad14000037b3</requestId>
			<serviceName>AdGroupService</serviceName>
			<methodName>get</methodName>
			<operations>1</operations>
			<responseTime>39</responseTime>
			<units>1</units>
		</ResponseHeader>
	</soap:Header>
	
	<soap:Body>
		<soap:Fault>
			<faultcode>soap:Server</faultcode>
			<faultstring>[RateExceededError &lt;rateName=RequestsPerMinute, rateKey=new_qps, rateScope=ACCOUNT, retryAfterSeconds=30&gt;]</faultstring>
			<detail>
			<ApiExceptionFault xmlns="https://adwords.google.com/api/adwords/cm/v201109">
			<message>[RateExceededError &lt;rateName=RequestsPerMinute, rateKey=new_qps, rateScope=ACCOUNT, retryAfterSeconds=30&gt;]</message>
			<ApplicationException.Type>ApiException</ApplicationException.Type>
			<errors xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:type="RateExceededError">
			<fieldPath></fieldPath>
			<trigger></trigger>
			<errorString>RateExceededError.RATE_EXCEEDED</errorString>
			<ApiError.Type>RateExceededError</ApiError.Type>
			<reason>RATE_EXCEEDED</reason>
			<rateName>RequestsPerMinute</rateName>
			<rateScope>ACCOUNT</rateScope>
			<retryAfterSeconds>30</retryAfterSeconds>
			</errors>
			</ApiExceptionFault>
			</detail>
		</soap:Fault>
	</soap:Body>
	</soap:Envelope>`
	
	decoder := xml.NewDecoder(bytes.NewBufferString(soapError))
	err := decoder.Decode(adgroupGet)
	if err != nil { panic(err) }
	
	// fmt.Printf("\nadgroupGet %+v\n", adgroupGet)
	ShouldEqual(`[RateExceededError <rateName=RequestsPerMinute, rateKey=new_qps, rateScope=ACCOUNT, retryAfterSeconds=30>]`, adgroupGet.Body.Fault.FaultString)
}