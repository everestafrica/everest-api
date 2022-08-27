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
	FindTransaction(transactionId string, userId string) (*models.AccountTransaction, error)
	FindAllTransactions(userId string, p types.Pagination) (*[]models.AccountTransaction, error)
	FindAllByInstitution(institution string, userId string, p types.Pagination) (*[]models.AccountTransaction, error)
	FindAllByType(txnType string, userId string, p types.Pagination) (*[]models.AccountTransaction, error)
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

func (r *accountTransactionRepo) FindTransaction(transactionId string, userId string) (*models.AccountTransaction, error) {
	var transaction models.AccountTransaction
	if err := r.db.Where("user_id = ? AND transaction_id", userId, transactionId).First(&transaction).Error; err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (r *accountTransactionRepo) FindAllTransactions(userId string, p types.Pagination) (*[]models.AccountTransaction, error) {
	var transactions []models.AccountTransaction
	if err := r.db.Scopes(paginate(p)).Where("user_id = ?", userId).Order("id DESC").Find(&transactions).Error; err != nil {
		return nil, err
	}

	return &transactions, nil
}

func (r *accountTransactionRepo) FindAllByInstitution(institution string, userId string, p types.Pagination) (*[]models.AccountTransaction, error) {
	var transactions []models.AccountTransaction
	if err := r.db.Scopes(paginate(p)).Where("user_id = ? AND institution = ?", userId, institution).Order("id DESC").Find(&transactions).Error; err != nil {
		return nil, err
	}

	return &transactions, nil
}
func (r *accountTransactionRepo) FindAllByType(txnType string, userId string, p types.Pagination) (*[]models.AccountTransaction, error) {
	var transactions []models.AccountTransaction
	if err := r.db.Scopes(paginate(p)).Where("user_id = ? AND type = ?", userId, txnType).Order("id DESC").Find(&transactions).Error; err != nil {
		return nil, err
	}

	return &transactions, nil
}

//chain := t.db.Scopes(paginate(types.Pagination, types.PageSize)).Preload("Counterparty").Where("user_id = ?", userId)
//
//if status != "" {
//chain = chain.Where("status = ?", status)
//}
//if entry != "" {
//chain = chain.Where("entry = ?", entry)
//}
//
//chain.Order("id DESC").Find(&transactions)
