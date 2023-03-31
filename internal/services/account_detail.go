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

type IAccountDetailsService interface {
	SetAccountDetails(code, userId string) error
	GetAccountDetails(accountId string) (*models.AccountDetail, error)
	GetAllAccountsDetails(userId string) (*[]models.AccountDetail, error)
	UnlinkAccount(id string) error
}

type accountDetailsService struct {
	userRepo           repositories.IUserRepository
	monoUserRepo       repositories.IMonoUserRepository
	accountDetailsRepo repositories.IAccountDetailsRepository
}

// NewAccountDetailsService will instantiate AccountDetailsService
func NewAccountDetailsService() IAccountDetailsService {
	return &accountDetailsService{
		userRepo:           repositories.NewUserRepo(),
		monoUserRepo:       repositories.NewMonoUserRepo(),
		accountDetailsRepo: repositories.NewAccountDetailsRepo(),
	}
}

func (ad accountDetailsService) SetAccountDetails(code, userId string) error {
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
		account := models.AccountDetail{
			UserId:          userId,
			AccountId:       monoId.Id,
			Institution:     details.Account.Institution.Name,
			InstitutionType: institutionType,
			AccountNumber:   details.Account.AccountNumber,
			Balance:         details.Account.Balance,
			Currency:        details.Account.Currency,
		}
		err = ad.accountDetailsRepo.Create(&account)
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

func (ad accountDetailsService) GetAccountDetails(accountId string) (*models.AccountDetail, error) {
	account, err := ad.accountDetailsRepo.FindByAccountId(accountId)
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (ad accountDetailsService) GetAllAccountsDetails(userId string) (*[]models.AccountDetail, error) {
	accounts, err := ad.accountDetailsRepo.FindAllByUserId(userId)
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

func (ad accountDetailsService) ReauthoriseUser(userId string) (*string, error) {
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

func (ad accountDetailsService) UnlinkAccount(id string) error {
	err := mono.Unlink(id)
	if err != nil {
		return err
	}
	err = ad.accountDetailsRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (ad accountDetailsService) GetInstitutionType(institutionName string) types.InstitutionType {
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
		return types.DepositAccount
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
