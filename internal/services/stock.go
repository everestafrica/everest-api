package services

import (
	"github.com/everestafrica/everest-api/internal/commons/log"
	"github.com/everestafrica/everest-api/internal/external/asset"
	"github.com/everestafrica/everest-api/internal/models"
	"github.com/everestafrica/everest-api/internal/repositories"
	"strings"
)

type IStockService interface {
	SetStockData() error
	DeleteStockData() error
}

type stockService struct {
	stockRepo repositories.IStockRepository
}

func NewStockService() IStockService {
	return &stockService{
		stockRepo: repositories.NewStockRepo(),
	}
}

func (s stockService) SetStockData() error {
	stocks, err := asset.ScrapeStockData()
	if err != nil {
		return err
	}
	stocksInDB, err := s.stockRepo.FindAllStockAssets()
	p, err := asset.GetAssetPrice("MSFT", false)
	if err != nil {
		return err
	}
	log.Info("p: ", p)
	for _, stock := range stocks {
		price, _ := asset.GetAssetPrice(strings.ToLower(stock.Symbol), false)
		log.Info("price: ", price)
		stk := &models.Stock{
			Name:   stock.Name,
			Image:  stock.Image,
			Symbol: stock.Symbol,
			//Price:  strconv.Itoa(int(*price)),
		}
		if len(*stocksInDB) < 1 {
			err = s.stockRepo.Create(stk)
			if err != nil {
				return err
			}
		} else {
			for _, v := range *stocksInDB {
				if v.Name != stock.Name {
					err = s.stockRepo.Create(stk)
					if err != nil {
						return err
					}
				} else {
					return nil
				}
			}
		}

	}
	log.Info("stored stocks in db!")
	return nil
}

func (s stockService) DeleteStockData() error {
	err := s.stockRepo.Delete()
	if err != nil {
		return err
	}
	log.Info("deleted stocks from db!")
	return nil
}
