package base

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

type Sedo struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
	Cookie   string `json:"cookie"`
	Page     string `json:"page"`
}

func (sedo *Sedo) Signup(email, password string) error {
	return nil
}

func (sedo *Sedo) Login(email, password string) error {
	return nil
}

func (sedo *Sedo) Logout() error {
	return nil
}

func (sedo *Sedo) GetToken(refreshToken ...string) error {
	return nil
}

// GetDomains has opts... string arg.
// opts[0] is keyword param,
// opts[1] is page param.
func (sedo *Sedo) GetDomains(opts ...string) (domains []string, err error) {
	var body map[string]interface{}
	param := `safe_search=2&synonyms=true&listing_type%5B%5D=1&listing_type%5B%5D=2&listing_type%5B%5D=3&listing_type%5B%5D=4&auction_group%5B%5D=62&auction_event%5B%5D=573&price_start=0&price_end=0&price_currency=3&traffic_start=0&traffic_end=0&number_of_words_min=1&number_of_words_max=0&len_min=1&len_max=0&special_characters%5B%5D=3&special_characters%5B%5D=1&special_characters%5B%5D=2&cat%5B%5D=0&cat%5B%5D=0&cat%5B%5D=0&type=0&special_inventory=4&kws=contains&age_min=0&age_max=0&keyword=` + url.QueryEscape(opts[0]) + `&page=` + opts[1] + `&rel=2&orderdirection=1&domainIds=&cc=&member=&v=0.1&o=json&m=search&f=requestSearch&pagesize=500&keywords_join=AND&loadListingFeatured=true&language=us`
	response, err := Request("POST", "https://sedo.com/service/common.php", strings.NewReader(param), &http.Header{
		"user-agent":   {"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/114.0"},
		"content-type": {"application/x-www-form-urlencoded"},
	})
	if err != nil {
		return nil, err
	}
	json.NewDecoder(response.Body).Decode(&body)
	for _, domain := range body["b"].(map[string]interface{})["general"].(map[string]interface{})["searchRequest"].(map[string]interface{})["resultList"].([]interface{}) {
		domains = append(domains, domain.(map[string]interface{})["2"].(string))
	}
	return domains, nil
}
