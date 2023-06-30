package base

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type Dynadot struct {
	Cookie string `json:"cookie"`
	Page   string `json:"page"`
}

func (dn *Dynadot) Signup(email, password string) error {
	return nil
}

func (dn *Dynadot) Login(email, password string) error {
	return nil
}

func (dn *Dynadot) Logout() error {
	return nil
}

func (dn *Dynadot) GetToken(refreshToken ...string) error {
	return nil
}

// GetDomains has opts... string arg.
// opts[0] is page_index param.
func (dn *Dynadot) GetDomains(opts ...string) (domains []string, err error) {
	response, err := Request("GET", "https://www.dynadot.com/market/auction?keyword=&order_by=3&page_size=100&page_index="+opts[0], nil, &http.Header{
		"user-agent": {"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/114.0"},
	})
	if err != nil {
		return nil, err
	}
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil, err
	}
	doc.Find(".link-domain-name").Each(func(i int, s *goquery.Selection) {
		domain, _ := s.Attr("title")
		domains = append(domains, domain)
	})
	return domains, nil
}
