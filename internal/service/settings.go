package service

import (
	"github.com/everestafrica/everest-api/internal/commons/types"
	"github.com/everestafrica/everest-api/internal/model"
	"github.com/everestafrica/everest-api/internal/repository"
)

type ISettingsService interface {
	CreateCustomCategory(category *types.CreateCustomCategory, userId string) error
	DeleteCustomCategory(userId string, categoryId string) error

	CreatePriceAlert(alert *types.CreatePriceAlert, userId string) error
	DeletePriceAlert(alertId string) error
	GetAllPriceAlerts(userId string) (*[]model.PriceAlert, error)

	CreateNewsInterest(interests *[]types.AddNewsInterest, userId string) error
	DeleteNewsInterest(interestId string) error
}

type settingsService struct {
	settingsRepo repository.ISettingsRepository
}

// NewSettingsService will instantiate SettingsService
func NewSettingsService() ISettingsService {
	return &settingsService{
		settingsRepo: repository.NewSettingsRepo(),
	}
}

func (s settingsService) CreateCustomCategory(category *types.CreateCustomCategory, userId string) error {
	c := model.Category{
		UserId: userId,
		Name:   category.Name,
		Icon:   category.Icon,
	}
	err := s.settingsRepo.CreateCustomCategory(&c)
	if err != nil {
		return err
	}
	return nil
}

func (s settingsService) DeleteCustomCategory(userId string, categoryId string) error {
	err := s.settingsRepo.DeleteCustomCategory(userId, categoryId)
	if err != nil {
		return err
	}
	return nil
}

func (s settingsService) CreatePriceAlert(alert *types.CreatePriceAlert, userId string) error {
	a := model.PriceAlert{
		UserId:   userId,
		Asset:    alert.Asset,
		IsCrypto: alert.IsCrypto,
		Target:   alert.Target,
	}
	err := s.settingsRepo.CreatePriceAlert(&a)
	if err != nil {
		return err
	}
	return nil
}

func (s settingsService) DeletePriceAlert(alertId string) error {
	err := s.settingsRepo.DeletePriceAlert(alertId)
	if err != nil {
		return err
	}
	return nil
}

func (s settingsService) GetAllPriceAlerts(userId string) (*[]model.PriceAlert, error) {
	alerts, err := s.settingsRepo.FindAllPriceAlerts(userId)
	if err != nil {
		return nil, err
	}
	return alerts, nil
}

func (s settingsService) CreateNewsInterest(interests *[]types.AddNewsInterest, userId string) error {
	for _, interest := range *interests {
		i := model.NewsInterest{
			UserId:   userId,
			Interest: types.NewsInterest(interest.Category),
		}
		err := s.settingsRepo.CreateNewsInterest(&i)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s settingsService) DeleteNewsInterest(interestId string) error {
	err := s.settingsRepo.DeleteNewsInterest(interestId)
	if err != nil {
		return err
	}
	return nil
}
