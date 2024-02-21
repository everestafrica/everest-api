package crypto

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/everestafrica/everest-api/internal/commons/types"
	"github.com/everestafrica/everest-api/internal/commons/utils"
	util "gorm.io/gorm/utils"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type rateResponse struct {
	Data []RateData `json:"data"`
}

type RateData struct {
	ID       string `json:"id"`
	Symbol   string `json:"symbol"`
	Name     string `json:"name"`
	PriceUsd string `json:"priceUsd"`
}

func GetRates() ([]RateData, error) {
	symbols := []string{"BTC", "SOL", "BNB", "ETH"}
	var rates rateResponse
	url := "https://api.coincap.io/v2/assets"
	v, err := Get(url, &rates)
	if err != nil {
		return nil, err
	}
	res := v.(*rateResponse)
	var ratesData []RateData
	for _, data := range res.Data {
		if util.Contains(symbols, data.Symbol) {
			ratesData = append(ratesData, data)
		}
	}
	return ratesData, nil
}

func GetHistoricalRate(symbol types.CoinSymbol, date string) (*float64, error) {
	coin := strings.ToLower(utils.GetCoinName(symbol))
	var rate HistoricalRate
	url := fmt.Sprintf("https://api.coingecko.com/api/v3/coins/%s/history?date=%s&localization=false", coin, date)
	resp, err := http.Get(url)

	if err != nil {
		//logger := log.WithField("error in Mono GET request", err)
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)
	if resp.StatusCode >= http.StatusBadRequest {
		return nil, errors.New(resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &rate)

	if err != nil {
		return nil, err
	}

	var value float64
	for k := range rate.MarketData.CurrentPrice {
		if k == "usd" {
			value = rate.MarketData.CurrentPrice[k]
		}
	}
	fmt.Println("value: ", value)

	return &value, nil
}

type HistoricalRate struct {
	ID     string `json:"id"`
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
	Image  struct {
		Thumb string `json:"thumb"`
		Small string `json:"small"`
	} `json:"image"`
	MarketData struct {
		CurrentPrice map[string]float64 `json:"current_price"`
		MarketCap    map[string]float64 `json:"market_cap"`
		TotalVolume  map[string]float64 `json:"total_volume"`
	} `json:"market_data"`
}
