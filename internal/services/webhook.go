package services

import (
	"github.com/everestafrica/everest-api/internal/commons/types"
	"github.com/everestafrica/everest-api/internal/repositories"
)

type IWebhookService interface {
}

type webhookService struct {
	userRepo       repositories.IUserRepository
	accountService IAccountService
}

// NewWebhookService will instantiate WebhookService
func NewWebhookService() IWebhookService {
	return &webhookService{
		userRepo:       repositories.NewUserRepo(),
		accountService: NewAccountService(),
	}
}

func (ws webhookService) MonoWebhook(payload types.MonoWebhookPayload) {

}
