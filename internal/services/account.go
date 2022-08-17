package services

import (
	"errors"
	"github.com/everestafrica/everest-api/internal/commons/types"
	"github.com/everestafrica/everest-api/internal/external/mono"
	"github.com/everestafrica/everest-api/internal/models"
	"github.com/everestafrica/everest-api/internal/repositories"
	"github.com/google/uuid"
)

type IAccountDetailsService interface {
	SetAccountDetails(code, userId string) error
	SetAccountTransactions(transaction *types.MonoTransactionResponse, userId string) error
}

type accountDetailsService struct {
	accountTransactionRepo repositories.IAccountTransactionRepository
	userRepo               repositories.IUserRepository
	accountDetailsRepo     repositories.IAccountDetailsRepository
}

// NewAccountDetailsService will instantiate AccountDetailsService
func NewAccountDetailsService() IAccountDetailsService {
	return &accountDetailsService{
		accountTransactionRepo: repositories.NewAccountTransactionRepo(),
		userRepo:               repositories.NewUserRepo(),
		accountDetailsRepo:     repositories.NewAccountDetailsRepo(),
	}
}

func (ad accountDetailsService) SetAccountDetails(code, userId string) error {
	monoCode := types.MonoAccountIdRequest{
		Code: code,
	}
	monoId, err := mono.GetAccountId(&monoCode)

	details, err := mono.GetAccountDetails(monoId.Id)
	if err != nil {
		return err
	}
	user, err := ad.accountDetailsRepo.FindByUserId(userId)
	if user != nil {
		user.Balance = details.Account.Balance
		err = ad.accountDetailsRepo.Update(user)
		if err != nil {
			return err
		}
	} else {
		if details.Meta.DataStatus == "AVAILABLE" {
			account := models.AccountDetail{
				UserId:        userId,
				AccountId:     uuid.NewString(),
				MonoId:        monoId.Id,
				Institution:   details.Account.Institution.Name,
				AccountNumber: details.Account.AccountNumber,
				Balance:       details.Account.Balance,
				Currency:      details.Account.Currency,
			}
			err := ad.accountDetailsRepo.Create(&account)
			if err != nil {
				return err
			}
		}
		if details.Meta.DataStatus == "FAILED" {
			return errors.New("request for user details failed")
		}
	}
	return nil
}

func (ad accountDetailsService) SetAccountTransactions(txn *types.MonoTransactionResponse, userId string) error {
	u, err := ad.userRepo.FindByUserId(userId)
	if err != nil {
		return err
	}

	for _, v := range txn.Data {
		transaction := models.AccountTransaction{
			UserId:        userId,
			MonoId:        u.MonoId,
			TransactionId: uuid.NewString(),
			//Institution:   "",
			Currency:  v.Currency,
			Amount:    v.Amount,
			Balance:   v.Balance,
			Date:      v.Date,
			Narration: v.Narration,
			Type:      types.TransactionType(v.Type),
			Category:  types.TransactionCategory(v.Category),
		}
		err := ad.accountTransactionRepo.Create(&transaction)
		if err != nil {
			return err
		}
	}

	return nil
}
