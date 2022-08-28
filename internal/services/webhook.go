package services

import (
	"errors"
	"github.com/everestafrica/everest-api/internal/commons/types"
	"github.com/everestafrica/everest-api/internal/models"
	"github.com/everestafrica/everest-api/internal/repositories"
	"github.com/google/uuid"
	"log"
)

type IWebhookService interface {
}

type webhookService struct {
	userRepo           repositories.IUserRepository
	accountDetailsRepo repositories.IAccountDetailsRepository
}

// NewWebhookService will instantiate WebhookService
func NewWebhookService() IWebhookService {
	return &webhookService{
		userRepo:           repositories.NewUserRepo(),
		accountDetailsRepo: repositories.NewAccountDetailsRepo(),
	}
}

func (ws webhookService) MonoWebhook(payload types.MonoWebhookPayload) error {
	if payload.Event == "mono.event.account_updated" {
		response := payload.Data.(types.MonoAccountResponse)
		if response.Meta.DataStatus == "AVAILABLE" {
			user, err := ws.userRepo.FindByMonoId(response.Account.ID)
			if err != nil {
				return errors.New("unable to find user with mono Id")
			}
			account := models.AccountDetail{
				UserId:        user.UserId,
				AccountId:     uuid.NewString(),
				MonoId:        response.Account.ID,
				Institution:   response.Account.Institution.Name,
				AccountNumber: response.Account.AccountNumber,
				Balance:       response.Account.Balance,
				Currency:      response.Account.Currency,
			}
			err = ws.accountDetailsRepo.Create(&account)
			if err != nil {
				return err
			}
		}
	}
	if payload.Event == "mono.events.account_connected" {
		response := payload.Data.(types.MonoWebhookConnect)
		log.Print(response)
		return nil
	}
	if payload.Event == "mono.events.account_unlinked" {
		response := payload.Data.(types.MonoWebhookUnlink)
		log.Print(response)
		return nil
	}
	if payload.Event == "mono.events.reauthorisation_required" {
		response := payload.Data.(types.MonoWebhookDataSync)
		log.Print(response)
		return nil
	}
	if payload.Event == "mono.events.account_reauthorized" {
		response := payload.Data.(types.MonoWebhookAccountSync)
		log.Print(response)
		return nil
	}
	if payload.Event == "mono.events.sync_failed" {
		response := payload.Data.(types.MonoWebhookDataSync)
		log.Print(response)
		return nil
	}
	if payload.Event == "mono.events.account_synced" {
		response := payload.Data.(types.MonoWebhookDataSync)
		log.Print(response)
		return nil
	}

	return nil
}
