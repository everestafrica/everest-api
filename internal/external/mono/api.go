package mono

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/everestafrica/everest-api/internal/config"
	"io"

	//"github.com/everestafrica/everest-api/internal/commons/log"
	"io/ioutil"
	"net/http"
)

var client = &http.Client{}
var baseEndpoint = "https://api.withmono.com"

func Post(url string, body, response interface{}) (interface{}, error) {
	requestJSON, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	r, _ := http.NewRequest(http.MethodPost, baseEndpoint+url, bytes.NewReader(requestJSON))
	r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", config.GetConf().MonoSecretKey))
	r.Header.Add("Content-Type", "application/json")

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
	if resp.StatusCode >= http.StatusBadRequest {
		return nil, errors.New(resp.Status)
	}

	b, err := ioutil.ReadAll(resp.Body)
	// s, _ := json.MarshalIndent(response, "", "\t")

	err = json.Unmarshal(b, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func Get(url string, query string, response interface{}) (interface{}, error) {
	r, _ := http.NewRequest(http.MethodGet, baseEndpoint+url+query, nil)
	r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", config.GetConf().MonoSecretKey))
	r.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(r)
	if err != nil {
		//logger := log.WithField("error in Mono GET request", err)
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)
	if resp.StatusCode >= http.StatusBadRequest {
		return nil, errors.New(resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	// s, _ := json.MarshalIndent(response, "", "\t")

	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
