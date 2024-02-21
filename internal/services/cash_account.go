package services

import (
	"errors"
	"github.com/everestafrica/everest-api/internal/commons/log"
	"github.com/everestafrica/everest-api/internal/commons/types"
	"github.com/everestafrica/everest-api/internal/external/mono"
	"github.com/everestafrica/everest-api/internal/models"
	"github.com/everestafrica/everest-api/internal/repositories"
	"time"
)

type ICashAccountService interface {
	SetCashAccountDetails(code, userId string) error
	GetCashAccountDetails(accountId string) (*models.CashAccount, error)
	GetAllCashAccountsDetails(userId string) (*[]models.CashAccount, error)
	UnlinkCashAccount(id string) error
}

type cashAccountService struct {
	userRepo     repositories.IUserRepository
	monoUserRepo repositories.IMonoUserRepository
	cashRepo     repositories.ICashAccountRepository
}

// NewCashAccountService will instantiate CashAccountService
func NewCashAccountService() ICashAccountService {
	return &cashAccountService{
		userRepo:     repositories.NewUserRepo(),
		monoUserRepo: repositories.NewMonoUserRepo(),
		cashRepo:     repositories.NewCashRepo(),
	}
}

func (ad cashAccountService) SetCashAccountDetails(code, userId string) error {
	monoCode := types.MonoAccountIdRequest{
		Code: code,
	}
	monoId, err := mono.GetAccountId(&monoCode)
	if err != nil {
		return err
	}

	monoUser := models.MonoUser{
		UserId:      userId,
		MonoId:      monoId.Id,
		MonoStatus:  "",
		LastRefresh: time.Now(),
		Reauth:      false,
		ReauthToken: nil,
	}
	err = ad.monoUserRepo.Create(&monoUser)
	if err != nil {
		return err
	}

	details, err := mono.GetAccountDetails(monoId.Id)
	if err != nil {
		return err
	}

	if details.Meta.DataStatus == "AVAILABLE" {
		institutionType := ad.GetInstitutionType(details.Account.Institution.Type)
		account := models.CashAccount{
			UserId:          userId,
			AccountId:       monoId.Id,
			Institution:     details.Account.Institution.Name,
			InstitutionType: institutionType,
			AccountNumber:   details.Account.AccountNumber,
			Balance:         details.Account.Balance,
			Currency:        details.Account.Currency,
		}
		err = ad.cashRepo.Create(&account)
		if err != nil {
			return err
		}
	}
	if details.Meta.DataStatus == "PROCESSING" {
		user, findErr := ad.monoUserRepo.FindByMonoId(monoId.Id)
		if findErr != nil {
			return findErr
		}
		user.MonoStatus = "PROCESSING"
		findErr = ad.monoUserRepo.Update(user)
		if findErr != nil {
			return findErr
		}

		// Send Email or Push Notification
		log.Info("processing user details collection")
		return errors.New("request for user details failed")
	}
	if details.Meta.DataStatus == "FAILED" {
		user, findErr := ad.monoUserRepo.FindByMonoId(monoId.Id)
		if findErr != nil {
			return findErr
		}
		user.MonoStatus = "FAILED"
		findErr = ad.monoUserRepo.Update(user)
		if findErr != nil {
			return findErr
		}

		// Send Email or Push Notification
		log.Error("failed to collect user details from institution")
		return errors.New("request for user details failed")
	}
	//}
	return nil
}

func (ad cashAccountService) GetCashAccountDetails(accountId string) (*models.CashAccount, error) {
	account, err := ad.cashRepo.FindByAccountId(accountId)
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (ad cashAccountService) GetAllCashAccountsDetails(userId string) (*[]models.CashAccount, error) {
	accounts, err := ad.cashRepo.FindAllByUserId(userId)
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

func (ad cashAccountService) ReauthoriseUser(userId string) (*string, error) {
	user, err := ad.monoUserRepo.FindByUserId(userId)
	result, err := mono.ReauthoriseUser(user.MonoId)
	if err != nil {
		return nil, err
	}
	user.ReauthToken = &result.Token
	err = ad.monoUserRepo.Update(user)
	if err != nil {
		return nil, err
	}
	return &result.Token, nil
}

func (ad cashAccountService) UnlinkCashAccount(id string) error {
	err := mono.Unlink(id)
	if err != nil {
		return err
	}
	err = ad.cashRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (ad cashAccountService) GetInstitutionType(institutionName string) types.InstitutionType {
	savings := []string{"Piggyvest", "Cowrywise"}
	wallets := []string{"Barter", "Wallets Africa"}
	investments := []string{"Risevest", "Trove", "Chaka"}

	switch {
	case contains(savings, institutionName):
		return types.SavingsAccount
	case contains(wallets, institutionName):
		return types.WalletAccount
	case contains(investments, institutionName):
		return types.InvestmentAccount
	default:
		return types.CashAccount
	}
}

func contains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}
