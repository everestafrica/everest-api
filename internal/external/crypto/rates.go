package crypto

import "gorm.io/gorm/utils"

type rateResponse struct {
	Data []struct {
		ID       string `json:"id"`
		Symbol   string `json:"symbol"`
		Name     string `json:"name"`
		PriceUsd string `json:"priceUsd"`
	} `json:"data"`
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
		if utils.Contains(symbols, data.Symbol) {
			ratesData = append(ratesData, data)
		}
	}
	return ratesData, nil
}
