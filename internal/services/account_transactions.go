package services

import (
	"errors"
	"fmt"
	"github.com/everestafrica/everest-api/internal/commons/types"
	"github.com/everestafrica/everest-api/internal/external/mono"
	"github.com/everestafrica/everest-api/internal/models"
	"github.com/everestafrica/everest-api/internal/repositories"
	"time"
)

type IAccountTransactionService interface {
	CreateManualTransaction(userId string, transaction *types.CreateTransactionRequest) error
	SetAccountTransactions(userId string) error
	GetTransaction(transactionId string) (*models.AccountTransaction, error)
	GetAllTransactions(userId string, pagination types.Pagination) (*[]models.AccountTransaction, error)
	GetInstitutionTransactions(institution string, userId string, pagination types.Pagination) (*[]models.AccountTransaction, error)
	GetTransactionsByType(txnType types.TransactionType, userId string, pagination types.Pagination) (*[]models.AccountTransaction, error)
	GetInflow(dateRange types.DateRange, userId string) (*types.TxnFlowResponse, error)
	GetOutflow(dateRange types.DateRange, userId string) (*types.TxnFlowResponse, error)
	DeleteManualTransaction(transactionId string) error
}

type accountTransactionService struct {
	userRepo               repositories.IUserRepository
	monoUserRepo           repositories.IMonoUserRepository
	accountDetailsRepo     repositories.IAccountDetailsRepository
	accountTransactionRepo repositories.IAccountTransactionRepository
}

// NewAccountTransactionService will instantiate AccountTransactionService
func NewAccountTransactionService() IAccountTransactionService {
	return &accountTransactionService{
		userRepo:               repositories.NewUserRepo(),
		monoUserRepo:           repositories.NewMonoUserRepo(),
		accountDetailsRepo:     repositories.NewAccountDetailsRepo(),
		accountTransactionRepo: repositories.NewAccountTransactionRepo(),
	}
}

func (ad accountTransactionService) CreateManualTransaction(userId string, transaction *types.CreateTransactionRequest) error {
	txn := models.AccountTransaction{
		UserId:        userId,
		AccountId:     &userId, // or nil -- I'm not really sure
		TransactionId: transaction.TransactionId,
		Institution:   "Everest",
		Currency:      transaction.Currency,
		Amount:        transaction.Amount,
		Balance:       nil,
		Narration:     transaction.Narration,
		Merchant:      transaction.Merchant,
		IsRecurring:   transaction.IsRecurring,
		Type:          transaction.Type,
		Category:      transaction.Category,
		Date:          transaction.Date,
	}
	err := ad.accountTransactionRepo.Create(&txn)
	if err != nil {
		return err
	}
	return nil
}

func (ad accountTransactionService) DeleteManualTransaction(transactionId string) error {
	err := ad.accountTransactionRepo.Delete(transactionId)
	if err != nil {
		return err
	}
	return nil
}

func (ad accountTransactionService) SetAccountTransactions(userId string) error {
	u, err := ad.monoUserRepo.FindByUserId(userId)
	if err != nil {
		return err
	}
	lastRefresh := u.LastRefresh
	refreshTimeLimit := time.Now().Add(-12 * time.Hour)
	if lastRefresh.Before(refreshTimeLimit) {
		return errors.New("unable to refresh transaction data at the moment")
	}

	account, err := ad.accountDetailsRepo.FindByAccountId(u.MonoId)
	if err != nil {
		return err
	}

	txn, err := mono.GetAccountTransactions(u.MonoId)
	if err != nil {
		return err
	}
	for _, v := range txn.Data {
		bal := float64(v.Balance / 100)
		transaction := models.AccountTransaction{
			UserId:        userId,
			AccountId:     &u.MonoId,
			TransactionId: v.ID,
			Institution:   account.Institution,
			Currency:      types.CurrencySymbol(v.Currency),
			Amount:        float64(v.Amount / 100),
			Balance:       &bal,
			Date:          v.Date,
			Narration:     v.Narration,
			Type:          types.TransactionType(v.Type),
			Category:      types.TransactionCategory(v.Category),
		}
		err = ad.accountTransactionRepo.Create(&transaction)
		if err != nil {
			return err
		}
	}

	u.LastRefresh = time.Now()
	err = ad.monoUserRepo.Update(u)
	if err != nil {
		return err
	}
	return nil
}

func (ad accountTransactionService) GetTransaction(transactionId string) (*models.AccountTransaction, error) {
	transactions, err := ad.accountTransactionRepo.FindTransaction(transactionId)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (ad accountTransactionService) GetAllTransactions(userId string, pagination types.Pagination) (*[]models.AccountTransaction, error) {
	transactions, err := ad.accountTransactionRepo.FindAllTransactions(userId, pagination)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (ad accountTransactionService) GetInstitutionTransactions(institution string, userId string, pagination types.Pagination) (*[]models.AccountTransaction, error) {
	transactions, err := ad.accountTransactionRepo.FindAllByInstitution(institution, userId, pagination)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (ad accountTransactionService) UpdateTransaction(transactionId string, req *types.UpdateTransactionRequest) error {
	transaction, err := ad.GetTransaction(transactionId)
	if err != nil {
		return err
	}
	transaction.Category = req.Category
	transaction.IsRecurring = req.IsRecurring
	err = ad.accountTransactionRepo.Update(transaction)
	if err != nil {
		return err
	}
	return nil
}

func (ad accountTransactionService) GetTransactionsByType(txnType types.TransactionType, userId string, pagination types.Pagination) (*[]models.AccountTransaction, error) {
	transactions, err := ad.accountTransactionRepo.FindAllByType(txnType, userId, pagination)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (ad accountTransactionService) GetInflow(dateRange types.DateRange, userId string) (*types.TxnFlowResponse, error) {
	transactions, err := ad.accountTransactionRepo.FindAllTxnFlow(types.Credit, dateRange, userId)
	if err != nil {
		return nil, err
	}
	var inflow float64
	for _, v := range *transactions {
		inflow += v.Amount
	}
	result := &types.TxnFlowResponse{
		Total:     inflow,
		DateRange: fmt.Sprintf("from: %s - to: %s", dateRange.From, dateRange.To),
	}
	return result, err
}
func (ad accountTransactionService) GetOutflow(dateRange types.DateRange, userId string) (*types.TxnFlowResponse, error) {
	transactions, err := ad.accountTransactionRepo.FindAllTxnFlow(types.Debit, dateRange, userId)
	if err != nil {
		return nil, err
	}
	var outflow float64
	for _, v := range *transactions {
		outflow += v.Amount

	}
	result := &types.TxnFlowResponse{
		Total:     outflow,
		DateRange: fmt.Sprintf("%s - %s", dateRange.From, dateRange.To),
	}
	return result, err
}

//@qq!92H230qk - coinmarket pwd
