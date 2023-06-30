package base

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type SeoDomains struct {
	Cookie string `json:"cookie"`
	Page   string `json:"page"`
}

func (sd *SeoDomains) Signup(email, password string) error {
	return nil
}

func (sd *SeoDomains) Login(email, password string) error {
	return nil
}

func (sd *SeoDomains) Logout() error {
	return nil
}

func (sd *SeoDomains) GetToken(refreshToken ...string) error {
	return nil
}

// GetDomains has opts... string arg.
// opts[0] is page param.
func (sd *SeoDomains) GetDomains(opts ...string) (domains []string, err error) {
	response, err := Request("GET", "https://list.seo.domains/?page="+opts[0], nil, &http.Header{
		"user-agent": {"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/114.0"},
		"cookie":     {"XSRF-TOKEN=eyJpdiI6IkRuYjZwMlRoTmNVaFQybEk4M0N1YWc9PSIsInZhbHVlIjoiYjhFSE5jOHRxZ2lycEpBWjNTc0J2eSs1dHBkbVlRd3NWWHRuZGJ2REhPcmo0T0NWeTFRN3FYQ3cyT2ZDb2pHdCIsIm1hYyI6IjYzOTg3OGJhNDNiZTA0YTU3YzlkMmQzN2UwNjg1MjlhZjdmYTJiYzFiZjFhMzM4NzRjMTFjZDllMjNkNWUxOGQifQ%3D%3D", `seodomains_session=eyJpdiI6Im5mSENrZTJqUlQrS1h6bXdVS1FmN3c9PSIsInZhbHVlIjoidmUrVWRUVlwvcHd4ZE5SOVhscWt1cHZYSHBTcll5UlpUK0NJaUlibzYwWFwvaWk3a1RacitzNHlFT3NKU0gxaWVDIiwibWFjIjoiMjlmYmNkNWM1ZDJhY2I2ZmZhNGNiOWY2NjliZmIwOTkxZDhlNGQ3YmIwMTEzYjhhOWIyYzBjM2E5ODE4OTRiMSJ9`},
	})
	if err != nil {
		return nil, err
	}
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil, err
	}
	doc.Find(".align-middle.font-weight-bold.domain-col > span").Each(func(i int, s *goquery.Selection) {
		domains = append(domains, s.Text())
	})
	return domains, nil
}
