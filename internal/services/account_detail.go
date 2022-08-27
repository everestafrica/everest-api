package services

import (
	"errors"
	"github.com/everestafrica/everest-api/internal/commons/types"
	"github.com/everestafrica/everest-api/internal/external/mono"
	"github.com/everestafrica/everest-api/internal/models"
	"github.com/everestafrica/everest-api/internal/repositories"
	"github.com/google/uuid"
	"log"
)

type IAccountDetailsService interface {
	SetAccountDetails(code, userId string) error
	GetAccountDetails(accountId string) (*models.AccountDetail, error)
	GetAllAccountsDetails(userId string) (*[]models.AccountDetail, error)
	UnlinkAccount(id string) error
}

type accountDetailsService struct {
	userRepo           repositories.IUserRepository
	accountDetailsRepo repositories.IAccountDetailsRepository
}

// NewAccountDetailsService will instantiate AccountDetailsService
func NewAccountDetailsService() IAccountDetailsService {
	return &accountDetailsService{
		userRepo:           repositories.NewUserRepo(),
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

	user, err := ad.userRepo.FindByUserId(userId)
	if err != nil {
		return err
	}

	user.MonoId = append(user.MonoId, monoId.Id)
	ad.userRepo.Update(user)

	details, err := mono.GetAccountDetails(monoId.Id)
	if err != nil {
		return err
	}

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
	if details.Meta.DataStatus == "PROCESSING" {
		log.Print("processing user details collection")
		return errors.New("request for user details failed")
	}
	if details.Meta.DataStatus == "FAILED" {
		log.Print("failed to collect user details from institution")
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
