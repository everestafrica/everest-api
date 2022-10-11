package services

import (
	"errors"
	"github.com/everestafrica/everest-api/internal/commons/log"
	"github.com/everestafrica/everest-api/internal/commons/types"
	"github.com/everestafrica/everest-api/internal/external/mono"
	"github.com/everestafrica/everest-api/internal/models"
	"github.com/everestafrica/everest-api/internal/repositories"
)

type IWebhookService interface {
	MonoWebhook(payload types.MonoWebhookPayload) error
}

type webhookService struct {
	monoUserRepo       repositories.IMonoUserRepository
	accountDetailsRepo repositories.IAccountDetailsRepository
}

// NewWebhookService will instantiate WebhookService
func NewWebhookService() IWebhookService {
	return &webhookService{
		monoUserRepo:       repositories.NewMonoUserRepo(),
		accountDetailsRepo: repositories.NewAccountDetailsRepo(),
	}
}

func (ws webhookService) MonoWebhook(payload types.MonoWebhookPayload) error {
	if payload.Event == "mono.event.account_updated" {
		response := payload.Data.(types.MonoAccountResponse)
		if response.Meta.DataStatus == "AVAILABLE" {
			user, err := ws.monoUserRepo.FindByMonoId(response.Account.ID)
			if err != nil {
				return errors.New("unable to find user with mono Id")
			}
			account := models.AccountDetail{
				UserId:        user.UserId,
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
		log.Info("response: ", response)
		return nil
	}
	if payload.Event == "mono.events.account_unlinked" {
		response := payload.Data.(types.MonoWebhookUnlink)
		user, err := ws.monoUserRepo.FindByMonoId(response.Account.Id)
		if err != nil {
			log.Error("mono id error", err)
			return errors.New("unable to find user with Mono Id")
		}
		err = ws.monoUserRepo.Delete(user.UserId)
		if err != nil {
			log.Error("delete mono user error", err)
			return err
		}
		return nil
	}
	if payload.Event == "mono.events.reauthorisation_required" {
		response := payload.Data.(types.MonoWebhookDataSync)
		user, err := ws.monoUserRepo.FindByMonoId(response.Account.Id)
		if err != nil {
			log.Error("mono id error", err)
			return errors.New("unable to find user with Mono Id")
		}
		user.Reauth = true
		err = ws.monoUserRepo.Update(user)
		if err != nil {
			return err
		}

		// send email or push notification to user
		result, err := mono.ReauthoriseUser(response.Account.Id)
		if err != nil {
			return err
		}
		user.ReauthToken = &result.Token
		err = ws.monoUserRepo.Update(user)
		if err != nil {
			return err
		}
		return nil
	}
	if payload.Event == "mono.events.account_reauthorized" {
		response := payload.Data.(types.MonoWebhookAccountSync)
		log.Info("response: ", response)
		return nil
	}
	if payload.Event == "mono.events.sync_failed" {
		response := payload.Data.(types.MonoWebhookDataSync)
		log.Info("response: ", response)
		return nil
	}
	if payload.Event == "mono.events.account_synced" {
		response := payload.Data.(types.MonoWebhookDataSync)
		log.Info("response: ", response)
		return nil
	}

	return nil
}
