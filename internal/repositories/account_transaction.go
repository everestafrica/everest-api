package repositories

import (
	"github.com/everestafrica/everest-api/internal/commons/types"
	"github.com/everestafrica/everest-api/internal/database"
	"github.com/everestafrica/everest-api/internal/models"
	"gorm.io/gorm"
)

type IAccountTransactionRepository interface {
	Create(transaction *models.AccountTransaction) error
	Update(transaction *models.AccountTransaction) error
	Delete(transactionId string) error
	FindTransaction(transactionId string) (*models.AccountTransaction, error)
	FindAllTransactions(userId string, p types.Pagination) (*[]models.AccountTransaction, error)
	FindAllByInstitution(institution string, userId string, p types.Pagination) (*[]models.AccountTransaction, error)
	FindAllByType(txnType types.TransactionType, userId string, p types.Pagination) (*[]models.AccountTransaction, error)
	FindAllTxnFlow(txnType types.TransactionType, dateRange types.DateRange, userId string) (*[]models.AccountTransaction, error)
}

type accountTransaction struct {
	db *gorm.DB
}

// NewAccountTransactionRepo  will instantiate AccountTransaction Repository
func NewAccountTransactionRepo() IAccountTransactionRepository {
	return &accountTransaction{
		db: database.DB(),
	}
}

func (r *accountTransaction) Create(transaction *models.AccountTransaction) error {
	return r.db.Create(transaction).Error
}

func (r *accountTransaction) Update(transaction *models.AccountTransaction) error {
	return r.db.Save(transaction).Error
}

func (r *accountTransaction) Delete(transactionId string) error {
	var transaction models.AccountTransaction
	return r.db.Where("transaction_id = ? AND institution = ?", transactionId, "Everest").Delete(&transaction).Error
}

func (r *accountTransaction) FindTransaction(transactionId string) (*models.AccountTransaction, error) {
	var transaction models.AccountTransaction
	if err := r.db.Where("transaction_id = ?", transactionId).First(&transaction).Error; err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (r *accountTransaction) FindAllTransactions(userId string, p types.Pagination) (*[]models.AccountTransaction, error) {
	var transactions []models.AccountTransaction
	if err := r.db.Scopes(paginate(p)).Where("user_id = ?", userId).Order("id DESC").Find(&transactions).Error; err != nil {
		return nil, err
	}

	return &transactions, nil
}

func (r *accountTransaction) FindAllByInstitution(institution string, userId string, p types.Pagination) (*[]models.AccountTransaction, error) {
	var transactions []models.AccountTransaction
	if err := r.db.Scopes(paginate(p)).Where("user_id = ? AND institution = ?", userId, institution).Order("id DESC").Find(&transactions).Error; err != nil {
		return nil, err
	}

	return &transactions, nil
}
func (r *accountTransaction) FindAllByType(txnType types.TransactionType, userId string, p types.Pagination) (*[]models.AccountTransaction, error) {
	var transactions []models.AccountTransaction
	if err := r.db.Scopes(paginate(p)).Where("user_id = ? AND type = ?", userId, txnType).Order("id DESC").Find(&transactions).Error; err != nil {
		return nil, err
	}

	return &transactions, nil
}

func (r *accountTransaction) FindAllTxnFlow(txnType types.TransactionType, dateRange types.DateRange, userId string) (*[]models.AccountTransaction, error) {
	var transactions []models.AccountTransaction
	if err := r.db.Where("user_id = ? AND type = ? AND date > ? AND date <= ?", userId, txnType, dateRange.From, dateRange.To).Order("id DESC").Find(&transactions).Error; err != nil {
		return nil, err
	}

	return &transactions, nil
}
