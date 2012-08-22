I didn't find a Go driver at https://developers.google.com/adwords/api/docs/clientlibraries so I decided to invent one myself.

This library is for query Google's ADX or Adwords services

Although this library is far from complete, what is here works.

##Installation

This library is go get-able

```sh
go get github.com/KarateCode/adx
```

##Example

Here's an example of how to query the CampaignService

```go  
package main
  
import (
  "github.com/KarateCode/adx"
  "fmt"
)

func main() {
  credentials := adx.Auth{
    Email: "your_email",
    Password: "your_password",
    ClientId: "the_client_id",
    DeveloperToken: "your_developer_token",
    Type: "line_item" // this field is not used internally, you may use it to attach any arbitrary data for your own purposes
    Version: "v201206", // currently only support v201206
    Sandbox: false,
  }
  
  adwords := adx.New(credentials)
  
  data := adx.CampaignGetSelector{
    Fields:   []string{"Id", "Name"}, 
    Field:    "Id", 
    Operator: "IN", 
    Values:   []string{"702011"},
  }
  campaignGet, err := adwords.CampaignService.Get(data)
  if err != nil {
    panic(err)
  }
  
  fmt.Printf("campaignGet: %+v", campaignGet)
}
```

Make sure to replace the various credentials with information from your own Adwords or ADX account


##Setup
For your convenience, in this library's "init" function, it will attempt to populate a variable called AdxPull and AdxPush from environmental variables.  This allows you to keep from accidentally committing sensitive information to a public git repository.

AdxPull is an object of type "adx.Auth".  It will populate from these environment variables:
AdxPushDeveloperToken
AdxPushClientId
AdxPushType
AdxPushEmail
AdxPushPassword
AdxPushVersion

To make sure these are set, I'd recommend adding lines like this to your .bashrc or .bash_profile:
  
```sh
export AdxPushEmail=youremail
export AdxPushPassword=yourpassword
export AdxPushVersion=v201206
export AdxPushSandbox=true
export AdxPushType=AdxItem
export AdxPushClientId=yourid
export AdxPushDeveloperToken=yourdevelopertoken
```

The AdxPush variable will populate from these environment variables:
AdxPullDeveloperToken
AdxPullClientId
AdxPullType
AdxPullEmail
AdxPullPassword
AdxPullVersion

Once all of this is set up, the example can now be reduced to this:
```go  
package main
  
import (
  "github.com/KarateCode/adx"
  "fmt"
)

func main() {
  adwords := adx.New(adx.AdxPush)
  
  data := adx.CampaignGetSelector{
    Fields:   []string{"Id", "Name"}, 
    Field:    "Id", 
    Operator: "IN", 
    Values:   []string{"702011"},
  }
  campaignGet, err := adwords.CampaignService.Get(data)
  if err != nil {
    panic(err)
  }
  
  fmt.Printf("campaignGet: %+v", campaignGet)
}
```
