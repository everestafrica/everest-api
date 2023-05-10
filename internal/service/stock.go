package service

import (
	"github.com/everestafrica/everest-api/internal/external/asset"
	"github.com/everestafrica/everest-api/internal/model"
	"github.com/everestafrica/everest-api/internal/repository"
	"log"
)

type IStockService interface {
	SetStockData() error
	DeleteStockData() error
}

type stockService struct {
	stockRepo repository.IStockRepository
}

func NewStockService() IStockService {
	return &stockService{
		stockRepo: repository.NewStockRepo(),
	}
}

func (s stockService) SetStockData() error {
	stocks, err := asset.ScrapeStockData()
	if err != nil {
		return err
	}
	stocksInDB, err := s.stockRepo.FindAllStockAssets()
	for _, stock := range stocks {
		stk := &model.Stock{
			Name:   stock.Name,
			Image:  stock.Image,
			Symbol: stock.Symbol,
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
	log.Println("stored stocks in db!")
	return nil
}

func (s stockService) DeleteStockData() error {
	err := s.stockRepo.Delete()
	if err != nil {
		return err
	}
	log.Println("deleted stocks from db!")
	return nil
}
