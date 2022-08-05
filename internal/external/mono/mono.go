package mono

import (
	"fmt"
	"github.com/everestafrica/everest-api/internal/config"
	"github.com/go-resty/resty/v2"
)

const baseUrl = "https://api.withmono.com"

func request() *resty.Request {
	client := resty.New().R()
	client.Header.Set("mono-sec-key", config.GetConf().MonoSecretKey)
	return client
}

func GetAccountId(code string) (*resty.Response, error) {
	url := fmt.Sprintf("%s/accounts/auth", baseUrl)
	resp, err := request().Get(url)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func GetAccountDetails(id string) {

}
func GetBalance() {
}
