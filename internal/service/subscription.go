package service

import (
	"github.com/everestafrica/everest-api/internal/commons/types"
	"github.com/everestafrica/everest-api/internal/commons/util"
	"github.com/everestafrica/everest-api/internal/model"
	"github.com/everestafrica/everest-api/internal/repository"
	"time"
)

type ISubscriptionService interface {
	AddSubscription(request *types.SubscriptionRequest, userId string) error
	GetSubscription(subId string, userId string) (*model.Subscription, error)
	GetAllSubscriptions(userId string) (*[]model.Subscription, error)
	DeleteSubscription(subId string, userId string) error
}

type subscriptionService struct {
	subscriptionRepo repository.ISubscriptionRepository
}

// NewSubscriptionService will instantiate SubscriptionService
func NewSubscriptionService() ISubscriptionService {
	return &subscriptionService{
		subscriptionRepo: repository.NewSubscriptionRepo(),
	}
}

func (ss subscriptionService) AddSubscription(request *types.SubscriptionRequest, userId string) error {
	nextPayment, _ := time.Parse("Jan 2, 2006 at 3:04pm (MST)", request.NextPayment)
	sub := model.Subscription{
		UserId:         userId,
		Product:        request.Product,
		Price:          request.Price,
		Currency:       request.Currency,
		Icon:           request.Icon,
		Frequency:      request.Frequency,
		NextPayment:    nextPayment,
		SubscriptionId: util.GetUUID(),
	}
	err := ss.subscriptionRepo.Create(&sub)
	if err != nil {
		return err
	}
	return nil
}

func (ss subscriptionService) GetAllSubscriptions(userId string) (*[]model.Subscription, error) {
	subs, err := ss.subscriptionRepo.FindAllByUserId(userId)
	if err != nil {
		return nil, err
	}
	return subs, nil
}

func (ss subscriptionService) GetSubscription(subId string, userId string) (*model.Subscription, error) {
	sub, err := ss.subscriptionRepo.FindByUserIdAndSubId(userId, subId)
	if err != nil {
		return nil, err
	}
	return sub, nil
}

func (ss subscriptionService) DeleteSubscription(subId string, userId string) error {
	err := ss.subscriptionRepo.Delete(userId, subId)
	if err != nil {
		return err
	}
	return nil
}
