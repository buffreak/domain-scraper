package base

import (
	"io"
	"log"
	"net/http"
	"time"

	"github.com/fatih/structs"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file: ", err)
	}
}

func Request(method string, url string, body io.Reader, headers *http.Header) (*http.Response, error) {
	var client *http.Client = &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Timeout: 30 * time.Second,
	}
	request, err := http.NewRequest(method, url, body)
	if headers != nil {
		request.Header = *headers
	}
	if err != nil {
		return nil, err
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func StructToMap(s interface{}) (map[string]interface{}, error) {
	return structs.Map(s), nil
}

func HandlePanic() {
	if r := recover(); r != nil {
		log.Printf("%+v\n", r)
	}
}
