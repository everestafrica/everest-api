package news

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/everestafrica/everest-api/internal/commons/utils"
	"github.com/everestafrica/everest-api/internal/config"
	"github.com/gocolly/colly"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

var client = &http.Client{}

type News struct {
	Title       string `json:"title"`
	Img         string `json:"img"`
	Author      string `json:"author"`
	Link        string `json:"link"`
	Date        string `json:"date"`
	Description string `json:"excerpt"`
}

type Response struct {
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
			fmt.Println(util.PrettyPrint(news))

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

func FetchNews() ([]*Response, error) {
	r, _ := http.NewRequest(http.MethodGet, config.GetConf().NewsApiUrl, nil)

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
	var responses []*Response
	err = json.Unmarshal(body, &responses)

	if err != nil {
		return nil, err
	}

	return responses, nil
}
