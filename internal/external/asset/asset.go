package asset

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/everestafrica/everest-api/internal/commons/log"
	"github.com/everestafrica/everest-api/internal/config"
	"github.com/everestafrica/everest-api/internal/external/crypto"
	"github.com/gocolly/colly"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

type assetResponse struct {
	QuoteResponse struct {
		Result []struct {
			Language                          string  `json:"language"`
			Region                            string  `json:"region"`
			QuoteType                         string  `json:"quoteType"`
			TypeDisp                          string  `json:"typeDisp"`
			QuoteSourceName                   string  `json:"quoteSourceName"`
			Triggerable                       bool    `json:"triggerable"`
			CustomPriceAlertConfidence        string  `json:"customPriceAlertConfidence"`
			Currency                          string  `json:"currency"`
			RegularMarketChangePercent        float64 `json:"regularMarketChangePercent"`
			RegularMarketPrice                float64 `json:"regularMarketPrice"`
			MessageBoardID                    string  `json:"messageBoardId"`
			ExchangeTimezoneName              string  `json:"exchangeTimezoneName"`
			ExchangeTimezoneShortName         string  `json:"exchangeTimezoneShortName"`
			GmtOffSetMilliseconds             int     `json:"gmtOffSetMilliseconds"`
			Market                            string  `json:"market"`
			EsgPopulated                      bool    `json:"esgPopulated"`
			Exchange                          string  `json:"exchange"`
			ShortName                         string  `json:"shortName"`
			LongName                          string  `json:"longName"`
			MarketState                       string  `json:"marketState"`
			Tradeable                         bool    `json:"tradeable"`
			CryptoTradeable                   bool    `json:"cryptoTradeable"`
			RegularMarketChange               float64 `json:"regularMarketChange"`
			RegularMarketTime                 int     `json:"regularMarketTime"`
			RegularMarketDayHigh              float64 `json:"regularMarketDayHigh"`
			RegularMarketDayRange             string  `json:"regularMarketDayRange"`
			RegularMarketDayLow               float64 `json:"regularMarketDayLow"`
			RegularMarketVolume               int     `json:"regularMarketVolume"`
			RegularMarketPreviousClose        float64 `json:"regularMarketPreviousClose"`
			Bid                               float64 `json:"bid"`
			Ask                               float64 `json:"ask"`
			BidSize                           int     `json:"bidSize"`
			AskSize                           int     `json:"askSize"`
			FullExchangeName                  string  `json:"fullExchangeName"`
			FinancialCurrency                 string  `json:"financialCurrency"`
			RegularMarketOpen                 float64 `json:"regularMarketOpen"`
			AverageDailyVolume3Month          int     `json:"averageDailyVolume3Month"`
			AverageDailyVolume10Day           int     `json:"averageDailyVolume10Day"`
			FiftyTwoWeekLowChange             float64 `json:"fiftyTwoWeekLowChange"`
			FiftyTwoWeekLowChangePercent      float64 `json:"fiftyTwoWeekLowChangePercent"`
			FiftyTwoWeekRange                 string  `json:"fiftyTwoWeekRange"`
			FiftyTwoWeekHighChange            float64 `json:"fiftyTwoWeekHighChange"`
			FiftyTwoWeekHighChangePercent     float64 `json:"fiftyTwoWeekHighChangePercent"`
			FiftyTwoWeekLow                   float64 `json:"fiftyTwoWeekLow"`
			FiftyTwoWeekHigh                  float64 `json:"fiftyTwoWeekHigh"`
			EarningsTimestamp                 int     `json:"earningsTimestamp"`
			EarningsTimestampStart            int     `json:"earningsTimestampStart"`
			EarningsTimestampEnd              int     `json:"earningsTimestampEnd"`
			TrailingAnnualDividendRate        float64 `json:"trailingAnnualDividendRate"`
			TrailingPE                        float64 `json:"trailingPE"`
			TrailingAnnualDividendYield       float64 `json:"trailingAnnualDividendYield"`
			EpsTrailingTwelveMonths           float64 `json:"epsTrailingTwelveMonths"`
			EpsForward                        float64 `json:"epsForward"`
			EpsCurrentYear                    float64 `json:"epsCurrentYear"`
			PriceEpsCurrentYear               float64 `json:"priceEpsCurrentYear"`
			SharesOutstanding                 int64   `json:"sharesOutstanding"`
			BookValue                         float64 `json:"bookValue"`
			FiftyDayAverage                   float64 `json:"fiftyDayAverage"`
			FiftyDayAverageChange             float64 `json:"fiftyDayAverageChange"`
			FiftyDayAverageChangePercent      float64 `json:"fiftyDayAverageChangePercent"`
			TwoHundredDayAverage              float64 `json:"twoHundredDayAverage"`
			TwoHundredDayAverageChange        float64 `json:"twoHundredDayAverageChange"`
			TwoHundredDayAverageChangePercent float64 `json:"twoHundredDayAverageChangePercent"`
			MarketCap                         int64   `json:"marketCap"`
			ForwardPE                         float64 `json:"forwardPE"`
			PriceToBook                       float64 `json:"priceToBook"`
			SourceInterval                    int     `json:"sourceInterval"`
			ExchangeDataDelayedBy             int     `json:"exchangeDataDelayedBy"`
			IpoExpectedDate                   string  `json:"ipoExpectedDate"`
			PrevName                          string  `json:"prevName"`
			NameChangeDate                    string  `json:"nameChangeDate"`
			AverageAnalystRating              string  `json:"averageAnalystRating"`
			FirstTradeDateMilliseconds        int64   `json:"firstTradeDateMilliseconds"`
			PriceHint                         int     `json:"priceHint"`
			DisplayName                       string  `json:"displayName"`
			Symbol                            string  `json:"symbol"`
		} `json:"result"`
		Error any `json:"error"`
	} `json:"quoteResponse"`
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

type hundredStock struct {
	Rank   string
	Name   string
	Symbol string
	Image  string
	Price  string
}

func GetAssetPrice(symbol string, isCrypto bool) (*float64, error) {
	if isCrypto {
		symbol = symbol + "-USD"
	}
	var asset assetResponse
	accessKey := config.GetConf().AssetAccessKey
	url := fmt.Sprintf("https://yfapi.net/v6/finance/quote?symbols=%s", symbol)
	r, _ := http.NewRequest(http.MethodGet, url, nil)
	r.Header.Add("X-API-KEY", accessKey)
	r.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(r)
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
	err = json.Unmarshal(body, &asset)
	if err != nil {
		return nil, err
	}
	if len(asset.QuoteResponse.Result) < 1 {
		return nil, errors.New(fmt.Sprintf("no price found for %s", symbol))
	}

	return &asset.QuoteResponse.Result[0].RegularMarketPrice, nil
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

func ScrapeStockData() ([]hundredStock, error) {
	base := "https://companiesmarketcap.com"
	countries, err := ScrapeCountriesUrl()
	if err != nil {
		return nil, err
	}
	var response []hundredStock
	for _, url := range countries {
		c := colly.NewCollector()

		var data hundredStock

		c.OnHTML(".marketcap-table > tbody", func(e *colly.HTMLElement) {
			e.ForEach("tr", func(i int, e *colly.HTMLElement) {
				image := e.ChildAttr(".name-td > .logo-container > img", "src")
				name := e.ChildText(".company-name")
				symbol := e.ChildText(".company-code")
				price := e.ChildText("td > .price")
				num := strconv.Itoa(i + 1)

				data = hundredStock{
					Rank:   num,
					Name:   name,
					Symbol: symbol,
					Image:  base + image,
					Price:  price,
				}
				response = append(response, data)
			})
		})
		//c.OnRequest(func(r *colly.Request) {
		//	log.Info("Visiting: ", r.URL.String())
		//})
		err = c.Visit(base + url)
		if err != nil {
			return nil, err
		}
	}
	return response, nil
}

func ScrapeCountriesUrl() ([]string, error) {
	c := colly.NewCollector()

	var response []string

	c.OnHTML(".marketcap-table > tbody", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(i int, e *colly.HTMLElement) {
			data := e.ChildAttr("td a", "href")
			response = append(response, data)
		})
	})
	c.OnRequest(func(r *colly.Request) {
		log.Info("Visiting: ", r.URL.String())
	})
	err := c.Visit("https://companiesmarketcap.com/all-countries/")
	if err != nil {
		return nil, err
	}
	return response, nil
}
