package repository

import (
	"github.com/everestafrica/everest-api/internal/database"
	"github.com/everestafrica/everest-api/internal/model"
	"gorm.io/gorm"
	"time"
)

type IStockRepository interface {
	FindAllStockAssets() (*[]model.Stock, error)
	Create(Stock *model.Stock) error
	Delete() error
}

type stockRepo struct {
	db *gorm.DB
}

// NewStockRepo will instantiate Stock Repository
func NewStockRepo() IStockRepository {
	return &stockRepo{
		db: database.DB(),
	}
}

func (r *stockRepo) FindAllStockAssets() (*[]model.Stock, error) {
	var stocks *[]model.Stock
	if err := r.db.Find(&stocks).Error; err != nil {
		return nil, err
	}
	return stocks, nil
}

func (r *stockRepo) Create(stock *model.Stock) error {
	return r.db.Create(stock).Error
}

func (r *stockRepo) Delete() error {
	var Stock model.Stock

	lastHalfHour := time.Now().Add(-30 * time.Minute)
	if err := r.db.Where("created_at < ?", lastHalfHour).Delete(&Stock).Error; err != nil {
		return err
	}
	return nil

}
