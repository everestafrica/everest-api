package repositories

import (
	"github.com/everestafrica/everest-api/internal/commons/types"
	"github.com/everestafrica/everest-api/internal/database"
	"github.com/everestafrica/everest-api/internal/models"
	"gorm.io/gorm"
)

type ICoinTransactionRepository interface {
	Create(transaction *models.CoinTransaction) error
	Update(transaction *models.CoinTransaction) error
	Delete(userId string, symbol types.CoinSymbol, address string) error
	FindByUserId(userId string) (*[]models.CoinTransaction, error)
	FindTransaction(hash string) (*models.CoinTransaction, error)
	FindAllTransactions(userId string, p types.Pagination) (*[]models.CoinTransaction, error)
	FindAllByTypeAndSymbol(txnType types.TransactionType, symbol string, userId string, p types.Pagination) (*[]models.CoinTransaction, error)
	FindAllTxnFlow(txnType types.TransactionType, dateRange types.DateRange, userId string) (*[]models.CoinTransaction, error)
}

type cryptoTransactionRepo struct {
	db *gorm.DB
}

// NewCryptoTransactionRepo  will instantiate CoinTransaction Repository
func NewCryptoTransactionRepo() ICoinTransactionRepository {
	return &cryptoTransactionRepo{
		db: database.DB(),
	}
}

func (r *cryptoTransactionRepo) Create(transaction *models.CoinTransaction) error {
	return r.db.Create(transaction).Error
}

func (r *cryptoTransactionRepo) Update(transaction *models.CoinTransaction) error {
	return r.db.Save(transaction).Error
}

func (r *cryptoTransactionRepo) Delete(userId string, symbol types.CoinSymbol, address string) error {
	var crypto models.CoinTransaction
	return r.db.Where("user_id = ? AND symbol = ? AND wallet_address = ?", userId, symbol, address).Delete(&crypto).Error
}

func (r *cryptoTransactionRepo) FindByUserId(userId string) (*[]models.CoinTransaction, error) {
	var transaction []models.CoinTransaction
	if err := r.db.Where("user_id = ?", userId).Find(&transaction).Error; err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (r *cryptoTransactionRepo) FindTransaction(hash string) (*models.CoinTransaction, error) {
	var transaction models.CoinTransaction
	if err := r.db.Where("hash = ?", hash).First(&transaction).Error; err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (r *cryptoTransactionRepo) FindAllTransactions(userId string, p types.Pagination) (*[]models.CoinTransaction, error) {
	var transactions []models.CoinTransaction
	if err := r.db.Scopes(paginate(p)).Where("user_id = ?", userId).Preload("Coin_wallets").Order("id DESC").Find(&transactions).Error; err != nil {
		return nil, err
	}

	return &transactions, nil
}

func (r *cryptoTransactionRepo) FindAllByTypeAndSymbol(txnType types.TransactionType, symbol string, userId string, p types.Pagination) (*[]models.CoinTransaction, error) {
	var transactions []models.CoinTransaction
	chain := r.db.Scopes(paginate(p)).Where("user_id = ?", userId).Order("id DESC").Find(&transactions)
	if txnType != "" {
		chain = chain.Where("type = ?", txnType)
	}
	if symbol != "" {
		chain = chain.Where("symbol = ?", symbol)
	}
	if err := chain.Error; err != nil {
		return nil, err
	}
	return &transactions, nil
}

func (r *cryptoTransactionRepo) FindAllTxnFlow(txnType types.TransactionType, dateRange types.DateRange, userId string) (*[]models.CoinTransaction, error) {
	var transactions []models.CoinTransaction
	if err := r.db.Where("user_id = ? AND type = ? AND date > ? AND date <= ?", userId, dateRange.From, dateRange.To, txnType).Order("id DESC").Find(&transactions).Error; err != nil {
		return nil, err
	}

	return &transactions, nil
}
