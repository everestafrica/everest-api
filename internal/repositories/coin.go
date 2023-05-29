package repositories

import (
	"github.com/everestafrica/everest-api/internal/commons/types"
	"github.com/everestafrica/everest-api/internal/database"
	"github.com/everestafrica/everest-api/internal/models"
	"gorm.io/gorm"
)

type ICoinRepository interface {
	Create(coin *models.CoinWallet) error
	Update(coin *models.CoinWallet) error
	Delete(userId string, symbol types.CoinSymbol, address string) error
	FindByUserId(userId string) (*[]models.CoinWallet, error)
	FindById(Id string) (*models.CoinWallet, error)
	FindByAddressAndSymbol(address string, symbol types.CoinSymbol) (*models.CoinWallet, error)
}

type coinRepo struct {
	db *gorm.DB
}

// NewCoinRepo will instantiate CoinDetails Repository
func NewCoinRepo() ICoinRepository {
	return &coinRepo{
		db: database.DB(),
	}
}

func (r *coinRepo) Create(coin *models.CoinWallet) error {
	return r.db.Create(coin).Error
}

func (r *coinRepo) Update(coin *models.CoinWallet) error {
	return r.db.Save(coin).Error
}

func (r *coinRepo) Delete(userId string, symbol types.CoinSymbol, address string) error {
	var coin models.CoinWallet
	return r.db.Where("user_id = ? AND symbol = ?  AND wallet_address = ?", userId, symbol, address).Delete(&coin).Error
}

func (r *coinRepo) FindByUserId(userId string) (*[]models.CoinWallet, error) {
	var coin []models.CoinWallet
	if err := r.db.Where("user_id = ?", userId).Find(&coin).Error; err != nil {
		return nil, err
	}

	return &coin, nil
}

func (r *coinRepo) FindByAddressAndSymbol(address string, symbol types.CoinSymbol) (*models.CoinWallet, error) {
	var coin models.CoinWallet
	if err := r.db.Where("wallet_address = ? AND symbol = ?", address, symbol).First(&coin).Error; err != nil {
		return nil, err
	}

	return &coin, nil
}
func (r *coinRepo) FindById(Id string) (*models.CoinWallet, error) {
	var coin models.CoinWallet
	if err := r.db.Where("id = ?", Id).First(&coin).Error; err != nil {
		return nil, err
	}

	return &coin, nil
}
