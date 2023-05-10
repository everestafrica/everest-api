package repository

import (
	"github.com/everestafrica/everest-api/internal/database"
	"github.com/everestafrica/everest-api/internal/model"
	"gorm.io/gorm"
	"time"
)

type INewsRepository interface {
	FindAllNews() (*[]model.News, error)
	Create(news *model.News) error
	Delete() error
}

type newsRepo struct {
	db *gorm.DB
}

// NewNewsRepo will instantiate News Repository
func NewNewsRepo() INewsRepository {
	return &newsRepo{
		db: database.DB(),
	}
}

func (r *newsRepo) FindAllNews() (*[]model.News, error) {
	var news *[]model.News
	if err := r.db.Find(&news).Error; err != nil {
		return nil, err
	}
	return news, nil
}

func (r *newsRepo) Create(news *model.News) error {
	return r.db.Create(news).Error
}

func (r *newsRepo) Delete() error {
	var news model.News

	lastHalfHour := time.Now().Add(-30 * time.Minute)
	if err := r.db.Where("created_at < ?", lastHalfHour).Delete(&news).Error; err != nil {
		return err
	}
	return nil

}
