package repositories

import (
	"github.com/everestafrica/everest-api/internal/database"
	"github.com/everestafrica/everest-api/internal/models"
	"gorm.io/gorm"
)

type INewsRepository interface {
	Create(news *models.News) error
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

func (r *newsRepo) Create(news *models.News) error {
	return r.db.Create(news).Error
}

func (r *newsRepo) Delete() error {
	var news models.News
	if err := r.db.Where("title = ?", "data").Delete(&news).Error; err != nil {
		return err
	}
	return nil

}
