package base

import (
	"encoding/json"
	"net/http"
	"strings"
)

type NameCheap struct {
	Cookie string `json:"cookie"`
	Page   string `json:"page"`
}

func (nc *NameCheap) Signup(email, password string) error {
	return nil
}

func (nc *NameCheap) Login(email, password string) error {
	return nil
}

func (nc *NameCheap) Logout() error {
	return nil
}

func (nc *NameCheap) GetToken(refreshToken ...string) error {
	return nil
}

// GetDomains has opts... string arg.
// opts[0] is page param.
func (nc *NameCheap) GetDomains(opts ...string) (domains []string, err error) {
	var body map[string]interface{}
	param := `{
		"operationName": "SaleTable",
		"variables": {
			"filter": {
				"saleType": "auction"
			},
			"sort": [{
				"column": "bidCount",
				"direction": "desc"
			}],
			"page": ` + opts[0] + `,
			"pageSize": 100
		},
		"extensions": {
			"persistedQuery": {
				"version": 1,
				"sha256Hash": "82f5b636090cafd054595e9d8e5172ee71bc9355741960efe443effea7fd8aaa"
			}
		}
	}`
	response, err := Request("POST", "https://aftermarketapi.namecheap.com/client/graphql", strings.NewReader(param), &http.Header{
		"user-agent":   {"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/114.0"},
		"content-type": {"application/json"},
	})
	if err != nil {
		return nil, err
	}
	json.NewDecoder(response.Body).Decode(&body)
	for _, domain := range body["data"].(map[string]interface{})["sales"].(map[string]interface{})["items"].([]interface{}) {
		domains = append(domains, domain.(map[string]interface{})["product"].(map[string]interface{})["name"].(string))
	}
	return domains, nil
}
