package repositories

import (
	"github.com/everestafrica/everest-api/internal/commons/types"
	"github.com/everestafrica/everest-api/internal/database"
	"github.com/everestafrica/everest-api/internal/models"
	"gorm.io/gorm"
)

type ICashTransactionRepository interface {
	Create(transaction *models.CashTransaction) error
	Update(transaction *models.CashTransaction) error
	Delete(transactionId string) error
	FindTransaction(transactionId string) (*models.CashTransaction, error)
	FindAllTransactions(userId string, p types.Pagination) (*[]models.CashTransaction, error)
	FindAllByInstitution(institution string, userId string, p types.Pagination) (*[]models.CashTransaction, error)
	FindAllByType(txnType types.TransactionType, userId string, p types.Pagination) (*[]models.CashTransaction, error)
	FindAllTxnFlow(txnType types.TransactionType, dateRange types.DateRange, userId string) (*[]models.CashTransaction, error)
}

type cashTransaction struct {
	db *gorm.DB
}

// NewCashTransactionRepo  will instantiate CashTransaction Repository
func NewCashTransactionRepo() ICashTransactionRepository {
	return &cashTransaction{
		db: database.DB(),
	}
}

func (r *cashTransaction) Create(transaction *models.CashTransaction) error {
	return r.db.Create(transaction).Error
}

func (r *cashTransaction) Update(transaction *models.CashTransaction) error {
	return r.db.Save(transaction).Error
}

func (r *cashTransaction) Delete(transactionId string) error {
	var transaction models.CashTransaction
	return r.db.Where("transaction_id = ? AND institution = ?", transactionId, "Everest").Delete(&transaction).Error
}

func (r *cashTransaction) FindTransaction(transactionId string) (*models.CashTransaction, error) {
	var transaction models.CashTransaction
	if err := r.db.Where("transaction_id = ?", transactionId).First(&transaction).Error; err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (r *cashTransaction) FindAllTransactions(userId string, p types.Pagination) (*[]models.CashTransaction, error) {
	var transactions []models.CashTransaction
	if err := r.db.Scopes(paginate(p)).Where("user_id = ?", userId).Order("id DESC").Find(&transactions).Error; err != nil {
		return nil, err
	}

	return &transactions, nil
}

func (r *cashTransaction) FindAllByInstitution(institution string, userId string, p types.Pagination) (*[]models.CashTransaction, error) {
	var transactions []models.CashTransaction
	if err := r.db.Scopes(paginate(p)).Where("user_id = ? AND institution = ?", userId, institution).Order("id DESC").Find(&transactions).Error; err != nil {
		return nil, err
	}

	return &transactions, nil
}
func (r *cashTransaction) FindAllByType(txnType types.TransactionType, userId string, p types.Pagination) (*[]models.CashTransaction, error) {
	var transactions []models.CashTransaction
	if err := r.db.Scopes(paginate(p)).Where("user_id = ? AND type = ?", userId, txnType).Order("id DESC").Find(&transactions).Error; err != nil {
		return nil, err
	}

	return &transactions, nil
}

func (r *cashTransaction) FindAllTxnFlow(txnType types.TransactionType, dateRange types.DateRange, userId string) (*[]models.CashTransaction, error) {
	var transactions []models.CashTransaction
	if err := r.db.Where("user_id = ? AND type = ? AND date > ? AND date <= ?", userId, txnType, dateRange.From, dateRange.To).Order("id DESC").Find(&transactions).Error; err != nil {
		return nil, err
	}

	return &transactions, nil
}
