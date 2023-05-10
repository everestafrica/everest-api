package repository

import (
	"github.com/everestafrica/everest-api/internal/commons/types"
	"github.com/everestafrica/everest-api/internal/database"
	"github.com/everestafrica/everest-api/internal/model"
	"gorm.io/gorm"
)

type IAssetRepository interface {
	Create(asset *model.Asset) error
	Update(asset *model.Asset) error
	Delete(userId string, symbol string) error
	FindByUserId(userId string) (*[]model.Asset, error)
	FindAsset(symbol string, userId string) (*model.Asset, error)
	FindAllAssets(userId string, p types.Pagination) (*[]model.Asset, error)
}

type assetRepo struct {
	db *gorm.DB
}

// NewAssetRepo  will instantiate Asset Repository
func NewAssetRepo() IAssetRepository {
	return &assetRepo{
		db: database.DB(),
	}
}

func (r *assetRepo) Create(asset *model.Asset) error {
	return r.db.Create(asset).Error
}

func (r *assetRepo) Update(asset *model.Asset) error {
	return r.db.Save(asset).Error
}

func (r *assetRepo) Delete(userId string, symbol string) error {
	var asset model.Asset
	return r.db.Where("user_id = ? AND symbol = ?", userId, symbol).Delete(&asset).Error
}

func (r *assetRepo) FindByUserId(userId string) (*[]model.Asset, error) {
	var asset []model.Asset
	if err := r.db.Where("user_id = ?", userId).Find(&asset).Error; err != nil {
		return nil, err
	}

	return &asset, nil
}

func (r *assetRepo) FindAsset(symbol string, userId string) (*model.Asset, error) {
	var asset model.Asset
	if err := r.db.Where("user_id = ? AND symbol", userId, symbol).First(&asset).Error; err != nil {
		return nil, err
	}

	return &asset, nil
}

func (r *assetRepo) FindAllAssets(userId string, p types.Pagination) (*[]model.Asset, error) {
	var assets []model.Asset
	if err := r.db.Scopes(paginate(p)).Where("user_id = ?", userId).Order("id DESC").Find(&assets).Error; err != nil {
		return nil, err
	}

	return &assets, nil
}
