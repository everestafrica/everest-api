package services

import (
	"errors"
	"github.com/everestafrica/everest-api/internal/commons/types"
	"github.com/everestafrica/everest-api/internal/external/mono"
	"github.com/everestafrica/everest-api/internal/models"
	"github.com/everestafrica/everest-api/internal/repositories"
	"github.com/google/uuid"
)

type IAccountService interface {
	SetAccountDetails(code, userId string) error
	SetAccountTransactions(transaction *types.MonoTransactionResponse, userId string) error
	GetAccountDetails(accountId string, userId string) (*models.AccountDetail, error)
	GetAllAccountsDetails(userId string) (*[]models.AccountDetail, error)
	UnlinkAccount(id string, userId string) error
}

type accountService struct {
	userRepo               repositories.IUserRepository
	accountTransactionRepo repositories.IAccountTransactionRepository
	accountDetailsRepo     repositories.IAccountDetailsRepository
}

// NewAccountService will instantiate AccountService
func NewAccountService() IAccountService {
	return &accountService{
		userRepo:               repositories.NewUserRepo(),
		accountTransactionRepo: repositories.NewAccountTransactionRepo(),
		accountDetailsRepo:     repositories.NewAccountDetailsRepo(),
	}
}

func (ad accountService) SetAccountDetails(code, userId string) error {
	monoCode := types.MonoAccountIdRequest{
		Code: code,
	}
	monoId, err := mono.GetAccountId(&monoCode)
	if err != nil {
		return err
	}

	details, err := mono.GetAccountDetails(monoId.Id)
	if err != nil {
		return err
	}

	user, err := ad.accountDetailsRepo.FindByUserId("", userId)
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

func (ad accountService) SetAccountTransactions(txn *types.MonoTransactionResponse, userId string) error {
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

func (ad accountService) GetAccountDetails(accountId string, userId string) (*models.AccountDetail, error) {
	account, err := ad.accountDetailsRepo.FindByUserId(accountId, userId)
	if err != nil {
		return nil, err
	}
	return account, nil
}
func (ad accountService) GetAllAccountsDetails(userId string) (*[]models.AccountDetail, error) {
	accounts, err := ad.accountDetailsRepo.FindAllByUserId(userId)
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

func (ad accountService) UnlinkAccount(id string, userId string) error {
	err := mono.Unlink(id)
	if err != nil {
		return err
	}
	return nil
}
