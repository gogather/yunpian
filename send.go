package yunpian

import (
	"github.com/gogather/com/http"
	"net/url"
)

type SMS struct {
	url        string
	apikey     string
	httpclient *http.HTTPClient
}

func NewSMS(url string, apikey string) *SMS {
	return &SMS{url: url, apikey: apikey, httpclient: http.NewHTTPClient()}
}

func (s *SMS) Send(phone string, text string) (string, error) {
	return s.httpclient.Post(s.url, url.Values{
		"apikey": {s.apikey},
		"mobile": {phone},
		"text":   {text},
	})
}
