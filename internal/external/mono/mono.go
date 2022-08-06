package mono

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/everestafrica/everest-api/internal/commons/types"
	"github.com/everestafrica/everest-api/internal/config"
	"io"
	"net/http"
)

type monoApi struct {
	baseURL   string
	secretKey string
}

const (
	accountAuth = "/v1/accounts/auth"
)

var client = &http.Client{}

func (api *monoApi) GetAccountId(request *types.MonoAccountIdRequest) (*string, error) {

	requestJSON, err := json.Marshal(request)

	if err != nil {
		return nil, err
	}

	r, _ := http.NewRequest(http.MethodPost,
		fmt.Sprintf("%s%s",
			api.baseURL, accountAuth), bytes.NewReader(requestJSON))

	r.Header.Add("mono-sec-key", config.GetConf().MonoSecretKey)
	r.Header.Add("Content-Type", "application/json")

	fmt.Println(r.URL)

	resp, err := client.Do(r)

	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusCreated {
		return nil, errors.New(resp.Status)
	}

	var res map[string]json.RawMessage
	err = json.NewDecoder(resp.Body).Decode(&res)

	if err != nil {
		return nil, err
	}

	var data map[string]string
	err = json.Unmarshal(res["data"], &data)

	id := data["id"]

	return &id, nil

}
