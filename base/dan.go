package base

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"

	"github.com/fatih/structs"
)

type Dan struct {
	Cookie string `json:"cookie"`
	Page   string `json:"page"`
}

func (dan *Dan) Signup(email, password string) error {
	return nil
}

func (dan *Dan) Login(email, password string) error {
	return nil
}

func (dan *Dan) Logout() error {
	return nil
}

func (dan *Dan) GetToken(refreshToken ...string) error {
	return nil
}

// GetDomains has opts... string arg.
// opts[0] is term param,
// opts[1] is offset param.
func (dan *Dan) GetDomains(opts ...string) (domains []string, err error) {
	var body map[string]interface{}
	response, err := Request("POST", `https://dan.com/search/search.json?&term=`+url.QueryEscape(opts[0])+`&order_by=recent-desc&offset=`+opts[1]+`&per_page=50`, strings.NewReader(""), &http.Header{
		"user-agent":   {"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/114.0"},
		"x-csrf-token": {"SECRYZVJREg2vwcQPvRLQhgLr48hCFgmX84bdzaJcb3NOcmrofHcz2iJw6B5qHpzWhZqeH_iVuEtvtDnjPPfmA"},
		"cookie":       {"visitor=vid=2d55fa53-601c-4971-856d-9380a5b55765", `_policy={"restricted_market":true,"tracking_market":"explicit"}`, `RT="z=1&dm=dan.com&si=9858cd1c-4592-4755-975f-a7fdab09bf3e&ss=ljinv6ex&sl=1&tt=180&rl=1&nu=uf6c1wh9&cl=6hr4"`, `cookie_permission=2024-06-30+13%3A18%3A36+%2B0200`, `intercom-id-e2frgfyw=2e5212ef-b8bf-46d9-aa5e-7d2befc0e971`, `intercom-session-e2frgfyw=`, `intercom-device-id-e2frgfyw=75c73daf-fae4-4d4e-a3e0-89a7a0c99708`, `OPTOUTMULTI=0:0%7Cc3:0%7Cc2:0%7Cc1:0%7Cc4:0`, `utag_main=v_id:01890c06c8f300a193ceece2a87005050001000d00978$_sn:2$_ss:0$_st:1688136591962$ses_id:1688134788963%3Bexp-session$_pn:2%3Bexp-session$isc:undefined%3Bexp-1688138388967`, `pwinteraction=Fri,%2030%20Jun%202023%2011:22:12%20GMT`, `bc9bd58fe1b6ef954d6d794db6d30e25e8ff50634d24346cf8006ef422e3c05c6e48b07678e34d08c97ad3f91012c80ac690b50f51fbd49b16e301de58d9c5ca=Bu1dkQNfKYMgwZhGLpw%2BjuKUqMtlzdIeEs8K%2FksRbAG9DFzxX5uXjMQZczx%2B2TBtpTvWjTUxs%2BRzQhY1sYjVMT573SBVcPy1PJj9J8uSaMo32D4bYPU1RR86cto7p1t9rgknDJOiSX%2F3STea441bqUWgpdIBQsIr7dfi%2Fg6W%2FX4h5%2BS6tGJhOeH4iMpeEv8VvdTo8XJZUruQRLiHYdxZreNiRx2%2BL9ZdG1GnzQhH18CZDEEjbbhH%2FUn8Kex4792TM9hsyYZIUjc8OxI7DEJi4yGu9hXLi5Qw7sV5MkVyg9QGkiAQXOhkOf9XGtDzbpf6rqgYJ5%2BhgJ01GAQxUgGSEjtqJ43SV3pm3xtof1hjXgdQl9OE96E3a3AGi1X9vLQ3SDexYxPyiS0c%2BmPnWEGdXOlcEqDBIuEm32L3%2FcRadQ%2Bn4EZbOhVenmEIl5KMu7PYymIorbRWl2Ni2Sw2PCIRSA%3D%3D--dc4mMTqwArkw1obh--r4d2i9uuDAMmCa2Vlz6Mjg%3D%3D`, `pathway=c2d3bc50-c4fd-4433-bed8-579f4597d197`, `ak_bmsc=9AAF843A03A23A6F268EF3F97A1476A0~000000000000000000000000000000~YAAQnqwwF8m4wwGJAQAATaWsDBSzBIGi2/WvAgImIvESd/fZW+EiCvB5VyRE3V/ETjsM4DvuOP6bkwm3hWLjfYGwu0u29H2uIITtF6jOW3U24thCBmtkKYTY/H5Bhg5XG5DLf+OSiP1jXNSFiSFLj2GO3NXyqUpPMnEtUHbdj57XyKFxLNjhy8u/i3o72BncGyGUGFjls/DJKBRujXgRluc9q5D/vA1wHRGEDqtZUSTxyI3sHsoQ46YPFJI/6d1k5juwmMBLqpoeozkbmj/5h4wETOT2oDoaOo+IpswuPjq401/ZeesUzlYGdjMeKHDrneN7xLlVr10xxjjQDrA+Tog4QXT1KJBINUigIMys3M+XnivEqGHSy93Y1j5GASD0jEdkPtr5L4E=`, `time_zone_offset=-420`, ` cookie_settings={%22performance%22:true%2C%22advertising%22:true%2C%22support%22:true}`, ` traffic=`, ` expBannerSplit=B`, `bm_sv=A623550AEA4ECE4282778F1A8E9B6296~YAAQrp7YF42ljQGJAQAAVHytDBT3I3+kniBNBXqVzZPG2nwxQrZQAB5HbWpbADoN+UbCM0E0YaJs8eAncHLgNai1NrNEMUeXv+Vntrn4w9yoD88pa59kLJpk+t/LJ9JH8NnXxEU0UtUjgu65hSvZiB1vYkF90E23/izSpKRaJmI1jfMIRBMloluhmi1kjcVQg30ZGgox+BtnJ2TpW93bGHabQcdXoDRuIdThULy+u2jxyoK6+myNqzhrPeXp~1`},
		"origin":       {"https://dan.com"},
	})
	if err != nil {
		return nil, err
	}
	json.NewDecoder(response.Body).Decode(&body)
	for _, domain := range body["domains"].([]interface{}) {
		domains = append(domains, domain.(map[string]interface{})["name"].(string))
	}
	return domains, nil
}

func (dan *Dan) StructToMap(s interface{}) (map[string]interface{}, error) {
	return structs.Map(s), nil
}
