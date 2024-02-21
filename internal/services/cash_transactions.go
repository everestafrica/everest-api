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

type ICashTransactionService interface {
	CreateManualTransaction(userId string, transaction *types.CreateTransactionRequest) error
	SetAccountTransactions(userId string) error
	GetTransaction(transactionId string) (*models.CashTransaction, error)
	GetAllTransactions(userId string, pagination types.Pagination) (*[]models.CashTransaction, error)
	GetInstitutionTransactions(institution string, userId string, pagination types.Pagination) (*[]models.CashTransaction, error)
	GetTransactionsByType(txnType types.TransactionType, userId string, pagination types.Pagination) (*[]models.CashTransaction, error)
	GetInflow(dateRange types.DateRange, userId string) (*types.TxnFlowResponse, error)
	GetOutflow(dateRange types.DateRange, userId string) (*types.TxnFlowResponse, error)
	DeleteManualTransaction(transactionId string) error
}

type cashTransactionService struct {
	userRepo            repositories.IUserRepository
	monoUserRepo        repositories.IMonoUserRepository
	cashRepo            repositories.ICashAccountRepository
	cashTransactionRepo repositories.ICashTransactionRepository
}

// NewAccountTransactionService will instantiate AccountTransactionService
func NewAccountTransactionService() ICashTransactionService {
	return &cashTransactionService{
		userRepo:            repositories.NewUserRepo(),
		monoUserRepo:        repositories.NewMonoUserRepo(),
		cashRepo:            repositories.NewCashRepo(),
		cashTransactionRepo: repositories.NewCashTransactionRepo(),
	}
}

func (cs cashTransactionService) CreateManualTransaction(userId string, transaction *types.CreateTransactionRequest) error {
	txn := models.CashTransaction{
		UserId:          userId,
		AccountId:       &userId, // or nil -- I'm not really sure
		Id:              transaction.TransactionId,
		Institution:     "Everest",
		InstitutionType: "Cash",
		Currency:        transaction.Currency,
		Amount:          transaction.Amount,
		Balance:         nil,
		Narration:       transaction.Narration,
		Merchant:        transaction.Merchant,
		IsRecurring:     transaction.IsRecurring,
		Type:            transaction.Type,
		Category:        transaction.Category,
		Date:            transaction.Date,
	}
	err := cs.cashTransactionRepo.Create(&txn)
	if err != nil {
		return err
	}
	return nil
}

func (cs cashTransactionService) DeleteManualTransaction(transactionId string) error {
	err := cs.cashTransactionRepo.Delete(transactionId)
	if err != nil {
		return err
	}
	return nil
}

func (cs cashTransactionService) SetAccountTransactions(userId string) error {
	u, err := cs.monoUserRepo.FindByUserId(userId)
	if err != nil {
		return err
	}
	lastRefresh := u.LastRefresh
	refreshTimeLimit := time.Now().Add(-12 * time.Hour)
	if lastRefresh.Before(refreshTimeLimit) {
		return errors.New("unable to refresh transaction data at the moment")
	}

	account, err := cs.cashRepo.FindByAccountId(u.MonoId)
	if err != nil {
		return err
	}

	txn, err := mono.GetAccountTransactions(u.MonoId)
	if err != nil {
		return err
	}
	for _, v := range txn.Data {
		bal := float64(v.Balance / 100)
		transaction := models.CashTransaction{
			UserId:      userId,
			AccountId:   &u.MonoId,
			Id:          v.ID,
			Institution: account.Institution,
			Currency:    types.CurrencySymbol(v.Currency),
			Amount:      float64(v.Amount / 100),
			Balance:     &bal,
			Date:        v.Date,
			Narration:   v.Narration,
			Type:        types.TransactionType(v.Type),
			Category:    types.TransactionCategory(v.Category),
		}
		err = cs.cashTransactionRepo.Create(&transaction)
		if err != nil {
			return err
		}
	}

	u.LastRefresh = time.Now()
	err = cs.monoUserRepo.Update(u)
	if err != nil {
		return err
	}
	return nil
}

func (cs cashTransactionService) GetTransaction(transactionId string) (*models.CashTransaction, error) {
	transactions, err := cs.cashTransactionRepo.FindTransaction(transactionId)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (cs cashTransactionService) GetAllTransactions(userId string, pagination types.Pagination) (*[]models.CashTransaction, error) {
	transactions, err := cs.cashTransactionRepo.FindAllTransactions(userId, pagination)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (cs cashTransactionService) GetInstitutionTransactions(institution string, userId string, pagination types.Pagination) (*[]models.CashTransaction, error) {
	transactions, err := cs.cashTransactionRepo.FindAllByInstitution(institution, userId, pagination)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (cs cashTransactionService) UpdateTransaction(transactionId string, req *types.UpdateTransactionRequest) error {
	transaction, err := cs.GetTransaction(transactionId)
	if err != nil {
		return err
	}
	transaction.Category = req.Category
	transaction.IsRecurring = req.IsRecurring
	err = cs.cashTransactionRepo.Update(transaction)
	if err != nil {
		return err
	}
	return nil
}

func (cs cashTransactionService) GetTransactionsByType(txnType types.TransactionType, userId string, pagination types.Pagination) (*[]models.CashTransaction, error) {
	transactions, err := cs.cashTransactionRepo.FindAllByType(txnType, userId, pagination)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (cs cashTransactionService) GetInflow(dateRange types.DateRange, userId string) (*types.TxnFlowResponse, error) {
	transactions, err := cs.cashTransactionRepo.FindAllTxnFlow(types.Credit, dateRange, userId)
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
func (cs cashTransactionService) GetOutflow(dateRange types.DateRange, userId string) (*types.TxnFlowResponse, error) {
	transactions, err := cs.cashTransactionRepo.FindAllTxnFlow(types.Debit, dateRange, userId)
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
