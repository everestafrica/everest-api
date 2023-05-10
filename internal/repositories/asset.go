package repositories

import (
	"github.com/everestafrica/everest-api/internal/commons/types"
	"github.com/everestafrica/everest-api/internal/database"
	"github.com/everestafrica/everest-api/internal/models"
	"gorm.io/gorm"
)

type IAssetRepository interface {
	Create(asset *models.Asset) error
	Update(asset *models.Asset) error
	Delete(userId string, symbol string) error
	FindByUserId(userId string) (*[]models.Asset, error)
	FindAsset(symbol string, userId string) (*models.Asset, error)
	FindAllAssets(userId string, p types.Pagination) (*[]models.Asset, error)
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

func (r *assetRepo) Create(asset *models.Asset) error {
	return r.db.Create(asset).Error
}

func (r *assetRepo) Update(asset *models.Asset) error {
	return r.db.Save(asset).Error
}

func (r *assetRepo) Delete(userId string, symbol string) error {
	var asset models.Asset
	return r.db.Where("user_id = ? AND symbol = ?", userId, symbol).Delete(&asset).Error
}

func (r *assetRepo) FindByUserId(userId string) (*[]models.Asset, error) {
	var asset []models.Asset
	if err := r.db.Where("user_id = ?", userId).Find(&asset).Error; err != nil {
		return nil, err
	}

	return &asset, nil
}

func (r *assetRepo) FindAsset(symbol string, userId string) (*models.Asset, error) {
	var asset models.Asset
	if err := r.db.Where("user_id = ? AND symbol", userId, symbol).First(&asset).Error; err != nil {
		return nil, err
	}

	return &asset, nil
}

func (r *assetRepo) FindAllAssets(userId string, p types.Pagination) (*[]models.Asset, error) {
	var assets []models.Asset
	if err := r.db.Scopes(paginate(p)).Where("user_id = ?", userId).Order("id DESC").Find(&assets).Error; err != nil {
		return nil, err
	}

	return &assets, nil
}
