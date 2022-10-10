package repositories

import (
	"github.com/everestafrica/everest-api/internal/commons/types"
	"github.com/everestafrica/everest-api/internal/database"
	"github.com/everestafrica/everest-api/internal/models"
	"gorm.io/gorm"
)

type ICryptoDetailsRepository interface {
	Create(crypto *models.CryptoDetail) error
	Update(crypto *models.CryptoDetail) error
	Delete(userId string, symbol types.CryptoSymbol, address string) error
	FindByUserId(userId string) (*[]models.CryptoDetail, error)
	FindByAddressAndSymbol(address string, symbol types.CryptoSymbol) (*models.CryptoDetail, error)
}

type cryptoDetailsRepo struct {
	db *gorm.DB
}

// NewCryptoDetailsRepo will instantiate CryptoDetails Repository
func NewCryptoDetailsRepo() ICryptoDetailsRepository {
	return &cryptoDetailsRepo{
		db: database.DB(),
	}
}

func (r *cryptoDetailsRepo) Create(crypto *models.CryptoDetail) error {
	return r.db.Create(crypto).Error
}

func (r *cryptoDetailsRepo) Update(crypto *models.CryptoDetail) error {
	return r.db.Save(crypto).Error
}

func (r *cryptoDetailsRepo) Delete(userId string, symbol types.CryptoSymbol, address string) error {
	var crypto models.CryptoDetail
	return r.db.Where("user_id = ? AND symbol = ?  AND wallet_address = ?", userId, symbol, address).Delete(&crypto).Error
}

func (r *cryptoDetailsRepo) FindByUserId(userId string) (*[]models.CryptoDetail, error) {
	var crypto []models.CryptoDetail
	if err := r.db.Where("user_id = ?", userId).Find(&crypto).Error; err != nil {
		return nil, err
	}

	return &crypto, nil
}

func (r *cryptoDetailsRepo) FindByAddressAndSymbol(address string, symbol types.CryptoSymbol) (*models.CryptoDetail, error) {
	var crypto models.CryptoDetail
	if err := r.db.Where("wallet_address = ? AND symbol = ?", address, symbol).First(&crypto).Error; err != nil {
		return nil, err
	}

	return &crypto, nil
}
