package repositories

import (
	"github.com/everestafrica/everest-api/internal/database"
	"github.com/everestafrica/everest-api/internal/models"
	"gorm.io/gorm"
	"time"
)

type IStockRepository interface {
	Create(Stock *models.Stock) error
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

func (r *stockRepo) Create(Stock *models.Stock) error {
	return r.db.Create(Stock).Error
}

func (r *stockRepo) Delete() error {
	var Stock models.Stock

	lastHalfHour := time.Now().Add(-30 * time.Minute)
	if err := r.db.Where("created_at < ?", lastHalfHour).Delete(&Stock).Error; err != nil {
		return err
	}
	return nil

}
