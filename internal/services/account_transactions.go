package services

import (
	"errors"
	"github.com/everestafrica/everest-api/internal/commons/types"
	"github.com/everestafrica/everest-api/internal/external/mono"
	"github.com/everestafrica/everest-api/internal/models"
	"github.com/everestafrica/everest-api/internal/repositories"
	"time"
)

type IAccountTransactionService interface {
	SetAccountTransactions(userId string) error
	GetTransaction(transactionId string, userId string) (*models.AccountTransaction, error)
	GetAllTransactions(userId string, pagination types.Pagination) (*[]models.AccountTransaction, error)
	GetInstitutionTransactions(institution string, userId string, pagination types.Pagination) (*[]models.AccountTransaction, error)
	GetTransactionsByType(txnType string, userId string, pagination types.Pagination) (*[]models.AccountTransaction, error)
}

type accountTransactionService struct {
	userRepo               repositories.IUserRepository
	accountDetailsRepo     repositories.IAccountDetailsRepository
	accountTransactionRepo repositories.IAccountTransactionRepository
}

// NewAccountTransactionService will instantiate AccountTransactionService
func NewAccountTransactionService() IAccountTransactionService {
	return &accountTransactionService{
		userRepo:               repositories.NewUserRepo(),
		accountDetailsRepo:     repositories.NewAccountDetailsRepo(),
		accountTransactionRepo: repositories.NewAccountTransactionRepo(),
	}
}

func (ad accountTransactionService) SetAccountTransactions(userId string) error {
	u, err := ad.userRepo.FindByUserId(userId)
	if err != nil {
		return err
	}
	lastRefresh := u.LastRefresh
	refreshTimeLimit := time.Now().Add(-12 * time.Hour)
	if lastRefresh.Before(refreshTimeLimit) {
		return errors.New("unable to refresh transaction data at the moment")
	}

	for _, monoId := range u.MonoId {
		account, err := ad.accountDetailsRepo.FindByAccountId(monoId)
		if err != nil {
			return err
		}

		txn, err := mono.GetAccountTransactions(monoId)
		if err != nil {
			return err
		}
		for _, v := range txn.Data {
			transaction := models.AccountTransaction{
				UserId:        userId,
				MonoId:        &monoId,
				TransactionId: v.ID,
				Institution:   account.Institution,
				Currency:      v.Currency,
				Amount:        float64(v.Amount / 100),
				Balance:       float64(v.Balance / 100),
				Date:          v.Date,
				Narration:     v.Narration,
				Type:          types.TransactionType(v.Type),
				Category:      types.TransactionCategory(v.Category),
			}
			err := ad.accountTransactionRepo.Create(&transaction)
			if err != nil {
				return err
			}
		}
	}
	u.LastRefresh = time.Now()
	ad.userRepo.Update(u)
	return nil
}

func (ad accountTransactionService) GetTransaction(transactionId string, userId string) (*models.AccountTransaction, error) {
	transactions, err := ad.accountTransactionRepo.FindTransaction(transactionId, userId)
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

func (ad accountTransactionService) GetTransactionsByType(txnType string, userId string, pagination types.Pagination) (*[]models.AccountTransaction, error) {
	transactions, err := ad.accountTransactionRepo.FindAllByType(txnType, userId, pagination)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}
