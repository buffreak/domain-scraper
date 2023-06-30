package base

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type DomainMarket struct {
	Cookie string `json:"cookie"`
	Page   string `json:"page"`
}

func (dm *DomainMarket) Signup(email, password string) error {
	return nil
}

func (dm *DomainMarket) Login(email, password string) error {
	return nil
}

func (dm *DomainMarket) Logout() error {
	return nil
}

func (dm *DomainMarket) GetToken(refreshToken ...string) error {
	return nil
}

// GetDomains has opts... string arg.
// opts[0] is searchInput param,
// opts[1] is page param.
func (dm *DomainMarket) GetDomains(opts ...string) (domains []string, err error) {
	response, err := Request("GET", "https://www.domainmarket.com/search?searchInput="+url.QueryEscape(opts[0])+"&searchType=c&excludedInput=&min_length=&max_length=&from_price=&to_price=&languageID=&tld%5B0%5D=all&sort=price&search=true&page="+opts[1], nil, &http.Header{
		"user-agent": {"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/114.0"},
	})
	if err != nil {
		return nil, err
	}
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil, err
	}
	doc.Find("a.url").Each(func(i int, s *goquery.Selection) {
		domain, _ := s.Attr("title")
		domains = append(domains, strings.ToLower(domain))
	})
	return domains, nil
}
