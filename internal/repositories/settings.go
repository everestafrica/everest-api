package repositories

import (
	"github.com/everestafrica/everest-api/internal/database"
	"github.com/everestafrica/everest-api/internal/models"
	"gorm.io/gorm"
)

type ISettingsRepository interface {
	CreateCustomCategory(category *models.Category) error
	DeleteCustomCategory(userId string, categoryId string) error
	FindAllCustomCategories(userId string) (*[]models.Category, error)
	CreatePriceAlert(alert *models.PriceAlert) error
	DeletePriceAlert(id string) error
	FindAllPriceAlerts(userId string) (*[]models.PriceAlert, error)
	CreateNewsInterest(interest *models.NewsInterest) error
	FindAllNewsInterests(userId string) (*[]models.NewsInterest, error)
	DeleteNewsInterest(id string) error
}

type settingsRepo struct {
	db *gorm.DB
}

// NewSettingsRepo will instantiate Settings Repository
func NewSettingsRepo() ISettingsRepository {
	return &settingsRepo{
		db: database.DB(),
	}
}

func (r *settingsRepo) CreateCustomCategory(category *models.Category) error {
	return r.db.Create(&category).Error
}

func (r *settingsRepo) FindAllCustomCategories(userId string) (*[]models.Category, error) {
	var categories []models.Category
	if err := r.db.Where("user_id = ? ", userId).Find(&categories).Error; err != nil {
		return nil, err
	}
	return &categories, nil
}

func (r *settingsRepo) DeleteCustomCategory(userId string, categoryId string) error {
	var category models.Category
	if err := r.db.Where("user_id = ? AND id = ? ", userId, categoryId).Delete(&category).Error; err != nil {
		return err
	}
	return nil
}

func (r *settingsRepo) CreatePriceAlert(alert *models.PriceAlert) error {
	return r.db.Create(&alert).Error
}
func (r *settingsRepo) DeletePriceAlert(id string) error {
	var alert models.PriceAlert
	if err := r.db.Where("id = ? ", id).Delete(&alert).Error; err != nil {
		return err
	}
	return nil
}

func (r *settingsRepo) FindAllPriceAlerts(userId string) (*[]models.PriceAlert, error) {
	var alerts []models.PriceAlert
	if err := r.db.Where("user_id = ? ", userId).Find(&alerts).Error; err != nil {
		return nil, err
	}
	return &alerts, nil
}

func (r *settingsRepo) UpdateAlert(alert models.PriceAlert) error {
	return r.db.Save(alert).Error
}

func (r *settingsRepo) CreateNewsInterest(interest *models.NewsInterest) error {
	return r.db.Create(&interest).Error
}

func (r *settingsRepo) FindAllNewsInterests(userId string) (*[]models.NewsInterest, error) {
	var interests []models.NewsInterest
	if err := r.db.Where("user_id = ? ", userId).Find(&interests).Error; err != nil {
		return nil, err
	}
	return &interests, nil
}

func (r *settingsRepo) DeleteNewsInterest(id string) error {
	var interest models.NewsInterest
	if err := r.db.Where("id = ? ", id).Delete(&interest).Error; err != nil {
		return err
	}
	return nil
}
