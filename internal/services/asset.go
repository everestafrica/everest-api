package services

import (
	"errors"
	"github.com/everestafrica/everest-api/internal/commons/types"
	"github.com/everestafrica/everest-api/internal/external/asset"
	"github.com/everestafrica/everest-api/internal/models"
	"github.com/everestafrica/everest-api/internal/repositories"
)

type IAssetService interface {
	AddAsset(symbol string, assetType string, userId string) error
	DeleteAsset(symbol string, userId string) error
	GetAsset(symbol string, userId string) (*models.Asset, error)
	GetAllAssets(userId string, pagination types.Pagination) (*[]models.Asset, error)
}

type assetService struct {
	userRepo  repositories.IUserRepository
	assetRepo repositories.IAssetRepository
}

func NewAssetService() IAssetService {
	return &assetService{
		userRepo:  repositories.NewUserRepo(),
		assetRepo: repositories.NewAssetRepo(),
	}
}
func (s assetService) AddAsset(symbol string, assetType string, userId string) error {
	var newAsset *models.Asset
	if assetType == "stock" {
		value, err := asset.GetCompanyStockValue(symbol)
		if err != nil {
			return err
		}
		name, err := asset.GetCompanyName(symbol)
		newAsset = &models.Asset{
			UserId: userId,
			Symbol: symbol,
			Name:   *name,
			Image:  "",
			Value:  *value,
		}
	}
	if assetType == "crypto" {
		return errors.New("no support for crypto assets yet")
	}
	err := s.assetRepo.Create(newAsset)
	if err != nil {
		return err
	}
	return nil
}

func (s assetService) DeleteAsset(symbol string, userId string) error {
	err := s.assetRepo.Delete(userId, symbol)
	if err != nil {
		return err
	}
	return nil
}

func (s assetService) GetAsset(symbol string, userId string) (*models.Asset, error) {
	a, err := s.assetRepo.FindAsset(symbol, userId)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (s assetService) GetAllAssets(userId string, pagination types.Pagination) (*[]models.Asset, error) {
	assets, err := s.assetRepo.FindAllAssets(userId, pagination)
	if err != nil {
		return nil, err
	}
	return assets, nil
}
