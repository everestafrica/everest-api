package services

import (
	"github.com/everestafrica/everest-api/internal/repositories"
)

type ISubscriptionService interface {
}

type subscriptionService struct {
	userRepo repositories.IUserRepository
}

// NewSubscriptionService will instantiate AccountDetailsService
func NewSubscriptionService() ISubscriptionService {
	return &subscriptionService{
		userRepo: repositories.NewUserRepo(),
	}
}
