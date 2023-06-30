package base

import (
	"io"
	"net/http"
	"regexp"
	"strings"

	"github.com/fatih/structs"
)

type GoDaddy struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
	Cookie   string `json:"cookie"`
	Baid     string `json:"baid"`
}

func (gd *GoDaddy) Signup(email, password string) error {
	return nil
}

func (gd *GoDaddy) Login(email, password string) error {
	return nil
}

func (gd *GoDaddy) Logout() error {
	return nil
}

func (gd *GoDaddy) GetToken(refreshToken ...string) error {
	return nil
}

// GetDomains has opts... string arg.
// opts[0] is t param,
// opts[1] is hidAdvSearch param,
// opts[2] is baid param.
func (gd *GoDaddy) GetDomains(opts ...string) (domains []string, err error) {
	param := `t=` + opts[0] + `&action=search&hidAdvSearch=` + opts[1] + `&rtr=7&baid=` + opts[2] + `&searchDir=1&isc=iddomgon1&JNRrxah=30a2620`
	response, err := Request("POST", "https://sg.auctions.godaddy.com/trpSearchResults.aspx", strings.NewReader(param), &http.Header{
		"user-agent":   {"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/114.0"},
		"content-type": {"application/x-www-form-urlencoded"},
	})
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	parseDomain := regexp.MustCompile(`View details for (.*?)"`).FindAllStringSubmatch(string(body), -1)
	for _, domain := range parseDomain {
		domains = append(domains, strings.TrimSpace(domain[1]))
	}
	gd.Baid = regexp.MustCompile(`s_baid1=(.*?);`).FindStringSubmatch(string(body))[1]
	return domains, nil
}

func (gd *GoDaddy) StructToMap(s interface{}) (map[string]interface{}, error) {
	return structs.Map(s), nil
}
