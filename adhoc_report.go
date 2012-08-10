package adx

import (
	"net/http"
	"net/url"
	"io"
	"bytes"
)

func (self *Conn) AdhocReport(reportString string) (io.Reader, error) {
	var params url.Values = make(url.Values)
	params.Add("__rdxml", reportString)
	req, err := http.NewRequest("POST", "https://adwords.google.com/api/adwords/reportdownload/" + self.Auth.Version, bytes.NewBufferString(params.Encode()))
	if err != nil {
		return nil, err
	}
	
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded") // VERY IMPORTANT. ADX wouldn't accept xml without it
	req.Header.Add("Authorization", "GoogleLogin auth=" + self.Token)
	req.Header.Add("clientCustomerId", self.Auth.ClientId)
	req.Header.Add("developerToken", self.Auth.DeveloperToken)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	return res.Body, nil
}