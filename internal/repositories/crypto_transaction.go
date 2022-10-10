package repositories

import (
	"github.com/everestafrica/everest-api/internal/commons/types"
	"github.com/everestafrica/everest-api/internal/database"
	"github.com/everestafrica/everest-api/internal/models"
	"gorm.io/gorm"
)

type ICryptoTransactionRepository interface {
	Create(transaction *models.CryptoTransaction) error
	Update(transaction *models.CryptoTransaction) error
	Delete(userId string, symbol types.CryptoSymbol, address string) error
	FindByUserId(userId string) (*[]models.CryptoTransaction, error)
}

type cryptoTransactionRepo struct {
	db *gorm.DB
}

// NewCryptoTransactionRepo  will instantiate CryptoTransaction Repository
func NewCryptoTransactionRepo() ICryptoTransactionRepository {
	return &cryptoTransactionRepo{
		db: database.DB(),
	}
}

func (r *cryptoTransactionRepo) Create(transaction *models.CryptoTransaction) error {
	return r.db.Create(transaction).Error
}

func (r *cryptoTransactionRepo) Update(transaction *models.CryptoTransaction) error {
	return r.db.Save(transaction).Error
}

func (r *cryptoTransactionRepo) Delete(userId string, symbol types.CryptoSymbol, address string) error {
	var crypto models.CryptoTransaction
	return r.db.Where("user_id = ? AND symbol = ? AND wallet_address = ?", userId, symbol, address).Delete(&crypto).Error
}

func (r *cryptoTransactionRepo) FindByUserId(userId string) (*[]models.CryptoTransaction, error) {
	var transaction []models.CryptoTransaction
	if err := r.db.Where("user_id = ?", userId).Find(&transaction).Error; err != nil {
		return nil, err
	}

	return &transaction, nil
}
