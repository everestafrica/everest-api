package services

import (
	"github.com/everestafrica/everest-api/internal/commons/types"
	"github.com/everestafrica/everest-api/internal/external/news"
	"github.com/everestafrica/everest-api/internal/models"
	"github.com/everestafrica/everest-api/internal/repositories"
	"log"
)

type INewsService interface {
	SetNews() error
	DeleteNews() error
	GetUserNews(userId string) (*[]models.News, error)
}

type newsService struct {
	newsRepo     repositories.INewsRepository
	settingsRepo repositories.ISettingsRepository
}

// NewNewsService will instantiate NewsService
func NewNewsService() INewsService {
	return &newsService{
		newsRepo:     repositories.NewNewsRepo(),
		settingsRepo: repositories.NewSettingsRepo(),
	}
}

func (ns newsService) GetUserNews(userId string) (*[]models.News, error) {
	userInterests, err := ns.settingsRepo.FindAllNewsInterests(userId)
	if err != nil {
		return nil, err
	}
	newsFromDB, err := ns.newsRepo.FindAllNews()
	if err != nil {
		return nil, err
	}
	interestsMap := make(map[types.NewsInterest]bool)
	for _, c := range *userInterests {
		interestsMap[c.Interest] = true
	}

	var filteredNews []models.News
	newsMap := make(map[string][]models.News)

	for _, n := range *newsFromDB {
		if interestsMap[types.NewsInterest(n.Category)] {
			newsMap[n.Category] = append(newsMap[n.Category], n)
		}
	}

	for _, c := range *userInterests {
		if allNews, ok := newsMap[string(c.Interest)]; ok {
			filteredNews = append(filteredNews, allNews...)
		}
	}

	return &filteredNews, nil
}

func (ns newsService) SetNews() error {
	scraped, err := news.ScrapeNews()
	fetched, err := news.FetchNews()
	if err != nil {
		return err
	}

	newsInDB, err := ns.newsRepo.FindAllNews()
	for _, s := range scraped {
		scrapedNews := models.News{
			Source:      "nairametrics",
			Title:       s.Title,
			Img:         s.Img,
			Author:      s.Author,
			Link:        s.Link,
			Description: s.Description,
			Date:        s.Date,
		}
		if len(*newsInDB) < 1 {
			err = ns.newsRepo.Create(&scrapedNews)
			if err != nil {
				return err
			}
		} else {
			for _, n := range *newsInDB {
				if n.Link != s.Link {
					err = ns.newsRepo.Create(&scrapedNews)
					if err != nil {
						return err
					}
				} else {
					return nil
				}
			}
		}

	}

	for _, f := range fetched.Data {
		fetchedNews := models.News{
			Source:      "newsapi",
			Title:       f.Title,
			Img:         f.UrlToImage,
			Author:      f.Author,
			Link:        f.Url,
			Description: f.Description,
			Date:        f.PublishedAt,
		}
		if len(*newsInDB) < 1 {
			err = ns.newsRepo.Create(&fetchedNews)
			if err != nil {
				return err
			}
		} else {
			for _, n := range *newsInDB {
				if n.Link != f.Url {
					err = ns.newsRepo.Create(&fetchedNews)
					if err != nil {
						return err
					}
				} else {
					return nil
				}
			}
		}
	}

	log.Print("stored in db!")
	return nil
}
func (ns newsService) DeleteNews() error {
	err := ns.newsRepo.Delete()
	if err != nil {
		return err
	}
	log.Print("deleted from db!")
	return nil
}
