package services

import (
	"github.com/everestafrica/everest-api/internal/commons/types"
	"github.com/everestafrica/everest-api/internal/models"
	"github.com/everestafrica/everest-api/internal/repositories"
)

type ISettingsService interface {
	CreateCustomCategory(category *types.CreateCustomCategory, userId string) error
	DeleteCustomCategory(categoryId string) error
	GetAllPriceAlerts(userId string) (*[]models.PriceAlert, error)
}

type settingsService struct {
	settingsRepo repositories.ISettingsRepository
}

// NewSettingsService will instantiate SettingsService
func NewSettingsService() ISettingsService {
	return &settingsService{
		settingsRepo: repositories.NewSettingsRepo(),
	}
}

func (s settingsService) CreateCustomCategory(category *types.CreateCustomCategory, userId string) error {
	c := models.CustomCategory{
		UserId: userId,
		Name:   category.Name,
		Emoji:  category.Emoji,
	}
	err := s.settingsRepo.CreateCustomCategory(&c)
	if err != nil {
		return err
	}
	return nil
}

func (s settingsService) DeleteCustomCategory(categoryId string) error {
	err := s.settingsRepo.DeleteCustomCategory(categoryId)
	if err != nil {
		return err
	}
	return nil
}

func (s settingsService) GetAllPriceAlerts(userId string) (*[]models.PriceAlert, error) {
	alerts, err := s.settingsRepo.FindAllPriceAlerts(userId)
	if err != nil {
		return nil, err
	}
	return alerts, nil
}
