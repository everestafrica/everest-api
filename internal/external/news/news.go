package news

import (
	"encoding/json"
	"errors"
	"github.com/everestafrica/everest-api/internal/config"
	"github.com/gocolly/colly"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type News struct {
	Title       string `json:"title"`
	Img         string `json:"img"`
	Author      string `json:"author"`
	Link        string `json:"link"`
	Date        string `json:"date"`
	Description string `json:"excerpt"`
}
type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    []Data `json:"data"`
}
type Data struct {
	Author      string `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	UrlToImage  string `json:"urlToImage"`
	PublishedAt string `json:"publishedAt"`
	Content     string `json:"content"`
}

func ScrapeNews() ([]*News, error) {
	c := colly.NewCollector()
	var news *News
	var response []*News

	c.OnHTML(".jeg_posts", func(e *colly.HTMLElement) {
		e.ForEach(".jeg_pl_lg_2", func(i int, e *colly.HTMLElement) {
			img := e.ChildAttr(".jeg_thumb > a > .thumbnail-container > img", "src")
			title := e.ChildText(".jeg_postblock_content > .jeg_post_title ")
			author := e.ChildText(".jeg_postblock_content > .jeg_post_meta .jeg_meta_author > a")
			date := e.ChildText(".jeg_postblock_content > .jeg_post_meta .jeg_meta_date > a ")
			excerpt := e.ChildText(".jeg_postblock_content > .jeg_post_excerpt > p ")
			link := e.ChildAttr(".jeg_post_title > a", "href")

			news = &News{
				Img:         img,
				Title:       title,
				Author:      author,
				Link:        link,
				Date:        date,
				Description: excerpt,
			}
			response = append(response, news)
		})

	})

	// Before making a request
	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL.String())
	})
	err := c.Visit("https://nairametrics.com/category/financial-literacy-for-nigerians/personal-finance/")
	if err != nil {
		return nil, err
	}
	return response, nil
}

func ScrapeCryptoNews() ([]*News, error) {
	c := colly.NewCollector()
	var news *News
	var response []*News

	c.OnHTML(".main", func(e *colly.HTMLElement) {
		e.ForEach(".card", func(i int, e *colly.HTMLElement) {
			link := e.ChildAttr("a", "href")
			img := e.ChildAttr("a > figure > img", "src")
			title := e.ChildText(".h-full > h3 > a ")
			date := e.ChildText(".text-gray > .date")
			news = &News{
				Img:         img,
				Title:       title,
				Author:      "",
				Link:        link,
				Date:        date,
				Description: "",
			}
			response = append(response, news)
		})

	})

	// Before making a request
	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL.String())
	})
	err := c.Visit("https://beincrypto.com/markets/")
	if err != nil {
		return nil, err
	}
	return response, nil
}

func FetchNews() (*Response, error) {
	resp, err := http.Get(config.GetConf().NewsApiUrl)
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
	var responses *Response
	err = json.Unmarshal(body, &responses)

	if err != nil {
		return nil, err
	}

	return responses, nil
}
