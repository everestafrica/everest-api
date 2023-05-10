package repository

import (
	"github.com/everestafrica/everest-api/internal/database"
	"github.com/everestafrica/everest-api/internal/model"
	"gorm.io/gorm"
)

type ISettingsRepository interface {
	CreateCustomCategory(category *model.Category) error
	DeleteCustomCategory(userId string, categoryId string) error
	FindAllCustomCategories(userId string) (*[]model.Category, error)
	CreatePriceAlert(alert *model.PriceAlert) error
	DeletePriceAlert(id string) error
	FindAllPriceAlerts(userId string) (*[]model.PriceAlert, error)
	CreateNewsInterest(interest *model.NewsInterest) error
	FindAllNewsInterests(userId string) (*[]model.NewsInterest, error)
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

func (r *settingsRepo) CreateCustomCategory(category *model.Category) error {
	return r.db.Create(&category).Error
}

func (r *settingsRepo) FindAllCustomCategories(userId string) (*[]model.Category, error) {
	var categories []model.Category
	if err := r.db.Where("user_id = ? ", userId).Find(&categories).Error; err != nil {
		return nil, err
	}
	return &categories, nil
}

func (r *settingsRepo) DeleteCustomCategory(userId string, categoryId string) error {
	var category model.Category
	if err := r.db.Where("user_id = ? AND id = ? ", userId, categoryId).Delete(&category).Error; err != nil {
		return err
	}
	return nil
}

func (r *settingsRepo) CreatePriceAlert(alert *model.PriceAlert) error {
	return r.db.Create(&alert).Error
}
func (r *settingsRepo) DeletePriceAlert(id string) error {
	var alert model.PriceAlert
	if err := r.db.Where("id = ? ", id).Delete(&alert).Error; err != nil {
		return err
	}
	return nil
}

func (r *settingsRepo) FindAllPriceAlerts(userId string) (*[]model.PriceAlert, error) {
	var alerts []model.PriceAlert
	if err := r.db.Where("user_id = ? ", userId).Find(&alerts).Error; err != nil {
		return nil, err
	}
	return &alerts, nil
}

func (r *settingsRepo) UpdateAlert(alert model.PriceAlert) error {
	return r.db.Save(alert).Error
}

func (r *settingsRepo) CreateNewsInterest(interest *model.NewsInterest) error {
	return r.db.Create(&interest).Error
}

func (r *settingsRepo) FindAllNewsInterests(userId string) (*[]model.NewsInterest, error) {
	var interests []model.NewsInterest
	if err := r.db.Where("user_id = ? ", userId).Find(&interests).Error; err != nil {
		return nil, err
	}
	return &interests, nil
}

func (r *settingsRepo) DeleteNewsInterest(id string) error {
	var interest model.NewsInterest
	if err := r.db.Where("id = ? ", id).Delete(&interest).Error; err != nil {
		return err
	}
	return nil
}
