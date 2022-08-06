package repositories

import (
	"github.com/everestafrica/everest-api/internal/database"
	"github.com/everestafrica/everest-api/internal/models"
	"gorm.io/gorm"
)

type IAccountTransactionRepository interface {
	Create(transaction *models.AccountTransaction) error
	Update(transaction *models.AccountTransaction) error
	FindByUserId(userId string) (*models.AccountTransaction, error)
}

type accountTransactionRepo struct {
	db *gorm.DB
}

// NewAccountTransactionRepo  will instantiate AccountTransaction Repository
func NewAccountTransactionRepo() IAccountTransactionRepository {
	return &accountTransactionRepo{
		db: database.DB(),
	}
}

func (r *accountTransactionRepo) Create(transaction *models.AccountTransaction) error {
	return r.db.Create(transaction).Error
}

func (r *accountTransactionRepo) Update(transaction *models.AccountTransaction) error {
	return r.db.Save(transaction).Error
}

func (r *accountTransactionRepo) FindByUserId(userId string) (*models.AccountTransaction, error) {
	var transaction models.AccountTransaction
	if err := r.db.Where("user_id = ?", userId).First(&transaction).Error; err != nil {
		return nil, err
	}

	return &transaction, nil
}
