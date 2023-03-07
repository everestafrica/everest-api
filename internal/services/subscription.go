package services

import (
	"github.com/everestafrica/everest-api/internal/commons/types"
	"github.com/everestafrica/everest-api/internal/models"
	"github.com/everestafrica/everest-api/internal/repositories"
	"time"
)

type ISubscriptionService interface {
	AddSubscription(request *types.SubscriptionRequest, userId string) error
	GetSubscription(subId int, userId string) (*models.Subscription, error)
	GetAllSubscriptions(userId string) (*[]models.Subscription, error)
	DeleteSubscription(subId int, userId string) error
}

type subscriptionService struct {
	subscriptionRepo repositories.ISubscriptionRepository
}

// NewSubscriptionService will instantiate SubscriptionService
func NewSubscriptionService() ISubscriptionService {
	return &subscriptionService{
		subscriptionRepo: repositories.NewSubscriptionRepo(),
	}
}

func (ss subscriptionService) AddSubscription(request *types.SubscriptionRequest, userId string) error {
	nextPayment, _ := time.Parse("Jan 2, 2006 at 3:04pm (MST)", request.NextPayment)
	sub := models.Subscription{
		UserId:      userId,
		Product:     request.Product,
		Price:       request.Price,
		Currency:    request.Currency,
		Icon:        request.Icon,
		Frequency:   request.Frequency,
		NextPayment: nextPayment,
	}
	err := ss.subscriptionRepo.Create(&sub)
	if err != nil {
		return err
	}
	return nil
}

func (ss subscriptionService) GetAllSubscriptions(userId string) (*[]models.Subscription, error) {
	subs, err := ss.subscriptionRepo.FindAllByUserId(userId)
	if err != nil {
		return nil, err
	}
	return subs, nil
}

func (ss subscriptionService) GetSubscription(subId int, userId string) (*models.Subscription, error) {
	sub, err := ss.subscriptionRepo.FindByUserIdAndSubId(userId, subId)
	if err != nil {
		return nil, err
	}
	return sub, nil
}

func (ss subscriptionService) DeleteSubscription(subId int, userId string) error {
	err := ss.subscriptionRepo.Delete(userId, subId)
	if err != nil {
		return err
	}
	return nil
}
