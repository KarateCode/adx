package v201309

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
)

func (self *Adwords) AdhocReport(reportString string) (io.Reader, error) {
	var params = make(url.Values)
	params.Add("__rdxml", reportString)
	req, err := http.NewRequest("POST", "https://adwords.google.com/api/adwords/reportdownload/"+self.CampaignService.Auth.Version, bytes.NewBufferString(params.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded") // VERY IMPORTANT. ADX wouldn't accept xml without it
	req.Header.Add("Authorization", "Bearer "+self.CampaignService.AccessToken)
	// req.Header.Add("Authorization", "GoogleLogin auth="+self.CampaignService.Token)
	req.Header.Add("clientCustomerId", self.CampaignService.Auth.ClientCustomerId)
	req.Header.Add("developerToken", self.CampaignService.Auth.DeveloperToken)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	return res.Body, nil
}
