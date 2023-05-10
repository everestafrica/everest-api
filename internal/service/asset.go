package service

import (
	"github.com/everestafrica/everest-api/internal/commons/types"
	"github.com/everestafrica/everest-api/internal/external/asset"
	"github.com/everestafrica/everest-api/internal/model"
	"github.com/everestafrica/everest-api/internal/repository"
)

type IAssetService interface {
	AddAsset(symbol string, isCrypto bool, userId string) error
	DeleteAsset(symbol string, userId string) error
	GetAsset(symbol string, userId string) (*model.Asset, error)
	GetAllAssets(userId string, pagination types.Pagination) (*[]model.Asset, error)
}

type assetService struct {
	userRepo  repository.IUserRepository
	assetRepo repository.IAssetRepository
}

func NewAssetService() IAssetService {
	return &assetService{
		userRepo:  repository.NewUserRepo(),
		assetRepo: repository.NewAssetRepo(),
	}
}
func (s assetService) AddAsset(symbol string, isCrypto bool, userId string) error {

	if !isCrypto {
		value, err := asset.GetAssetPrice(symbol, false)
		if err != nil {
			return err
		}
		name, err := asset.GetCompanyName(symbol)
		newAsset := &model.Asset{
			UserId: userId,
			Symbol: symbol,
			Name:   *name,
			Image:  "",
			Value:  *value,
		}
		err = s.assetRepo.Create(newAsset)
		if err != nil {
			return err
		}
		return nil
	} else {
		value, err := asset.GetAssetPrice(symbol, true)
		if err != nil {
			return err
		}
		name := asset.GetCoinName(types.CryptoSymbol(symbol))
		newAsset := &model.Asset{
			UserId: userId,
			Symbol: symbol,
			Name:   name,
			Image:  "",
			Value:  *value,
		}
		err = s.assetRepo.Create(newAsset)
		if err != nil {
			return err
		}
		return nil
	}

}

func (s assetService) DeleteAsset(symbol string, userId string) error {
	err := s.assetRepo.Delete(userId, symbol)
	if err != nil {
		return err
	}
	return nil
}

func (s assetService) GetAsset(symbol string, userId string) (*model.Asset, error) {
	a, err := s.assetRepo.FindAsset(symbol, userId)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (s assetService) GetAllAssets(userId string, pagination types.Pagination) (*[]model.Asset, error) {
	assets, err := s.assetRepo.FindAllAssets(userId, pagination)
	if err != nil {
		return nil, err
	}
	return assets, nil
}
