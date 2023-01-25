package asset

import (
	"fmt"
	"github.com/everestafrica/everest-api/internal/config"
	"github.com/everestafrica/everest-api/internal/external/crypto"
	"time"
)

type stockResponse struct {
	Pagination struct {
		Limit  int `json:"limit"`
		Offset int `json:"offset"`
		Count  int `json:"count"`
		Total  int `json:"total"`
	} `json:"pagination"`
	Data []struct {
		Open     float64 `json:"open"`
		High     float64 `json:"high"`
		Low      float64 `json:"low"`
		Last     float64 `json:"last"`
		Close    float64 `json:"close"`
		Volume   float64 `json:"volume"`
		Date     string  `json:"date"`
		Symbol   string  `json:"symbol"`
		Exchange string  `json:"exchange"`
	} `json:"data"`
}

type companyName struct {
	Pagination struct {
		Limit  int `json:"limit"`
		Offset int `json:"offset"`
		Count  int `json:"count"`
		Total  int `json:"total"`
	} `json:"pagination"`
	Data []struct {
		Name          string      `json:"name"`
		Symbol        string      `json:"symbol"`
		HasIntraday   bool        `json:"has_intraday"`
		HasEod        bool        `json:"has_eod"`
		Country       interface{} `json:"country"`
		StockExchange struct {
			Name        string `json:"name"`
			Acronym     string `json:"acronym"`
			Mic         string `json:"mic"`
			Country     string `json:"country"`
			CountryCode string `json:"country_code"`
			City        string `json:"city"`
			Website     string `json:"website"`
		} `json:"stock_exchange"`
	} `json:"data"`
}

type cryptoStat struct {
	Status struct {
		Timestamp    time.Time   `json:"timestamp"`
		ErrorCode    int         `json:"error_code"`
		ErrorMessage interface{} `json:"error_message"`
		Elapsed      int         `json:"elapsed"`
		CreditCount  int         `json:"credit_count"`
		Notice       interface{} `json:"notice"`
	} `json:"status"`
	Data struct {
		Sol []struct {
			ID             int       `json:"id"`
			Name           string    `json:"name"`
			Symbol         string    `json:"symbol"`
			Slug           string    `json:"slug"`
			NumMarketPairs int       `json:"num_market_pairs"`
			DateAdded      time.Time `json:"date_added"`
			Tags           []struct {
				Slug     string `json:"slug"`
				Name     string `json:"name"`
				Category string `json:"category"`
			} `json:"tags"`
			MaxSupply                     interface{} `json:"max_supply"`
			CirculatingSupply             float64     `json:"circulating_supply"`
			TotalSupply                   float64     `json:"total_supply"`
			IsActive                      int         `json:"is_active"`
			Platform                      interface{} `json:"platform"`
			CmcRank                       int         `json:"cmc_rank"`
			IsFiat                        int         `json:"is_fiat"`
			SelfReportedCirculatingSupply interface{} `json:"self_reported_circulating_supply"`
			SelfReportedMarketCap         interface{} `json:"self_reported_market_cap"`
			TvlRatio                      interface{} `json:"tvl_ratio"`
			LastUpdated                   time.Time   `json:"last_updated"`
			Quote                         struct {
				Usd struct {
					Price                 float64     `json:"price"`
					Volume24H             float64     `json:"volume_24h"`
					VolumeChange24H       float64     `json:"volume_change_24h"`
					PercentChange1H       float64     `json:"percent_change_1h"`
					PercentChange24H      float64     `json:"percent_change_24h"`
					PercentChange7D       float64     `json:"percent_change_7d"`
					PercentChange30D      float64     `json:"percent_change_30d"`
					PercentChange60D      float64     `json:"percent_change_60d"`
					PercentChange90D      float64     `json:"percent_change_90d"`
					MarketCap             float64     `json:"market_cap"`
					MarketCapDominance    float64     `json:"market_cap_dominance"`
					FullyDilutedMarketCap float64     `json:"fully_diluted_market_cap"`
					Tvl                   interface{} `json:"tvl"`
					LastUpdated           time.Time   `json:"last_updated"`
				} `json:"USD"`
			} `json:"quote"`
		} `json:"SOL"`
	} `json:"data"`
}

func GetCompanyStockValue(symbol string) (*float64, error) {
	var stock stockResponse
	accessKey := config.GetConf().StockAccessKey
	url := fmt.Sprintf("http://api.marketstack.com/v1/eod?access_key=%s&symbols=%s", accessKey, symbol)
	v, err := crypto.Get(url, &stock)
	if err != nil {
		return nil, err
	}
	res := v.(stockResponse)
	return &res.Data[0].Open, nil
}

func GetCompanyName(symbol string) (*string, error) {
	var name companyName
	accessKey := config.GetConf().StockAccessKey
	url := fmt.Sprintf("http://api.marketstack.com/v1/tickers?access_key=%s&symbols=%s", accessKey, symbol)
	v, err := crypto.Get(url, &name)
	if err != nil {
		return nil, err
	}
	res := v.(companyName)
	return &res.Data[0].Name, nil
}
func GetCryptoValue(symbol string) (*float64, error) {
	var stat cryptoStat
	accessKey := config.GetConf().CmcKey
	url := fmt.Sprintf("https://pro-api.coinmarketcap.com/v2/cryptocurrency/quotes/latest?CMC_PRO_API_KEY=%s&symbol=%s", accessKey, symbol)
	v, err := crypto.Get(url, &stat)
	if err != nil {
		return nil, err
	}
	res := v.(cryptoStat)
	return &res.Data.Sol[0].Quote.Usd.Price, nil
}

// XNAS XNYS
