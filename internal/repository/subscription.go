package repository

import (
	"github.com/everestafrica/everest-api/internal/database"
	"github.com/everestafrica/everest-api/internal/model"
	"gorm.io/gorm"
)

type ISubscriptionRepository interface {
	Create(sub *model.Subscription) error
	Update(sub *model.Subscription) error
	Delete(userId string, subscriptionId string) error
	FindAllByUserId(userId string) (*[]model.Subscription, error)
	FindByUserIdAndSubId(userId string, subscriptionId string) (*model.Subscription, error)
}

type subscriptionRepo struct {
	db *gorm.DB
}

// NewSubscriptionRepo will instantiate Subscription Repository
func NewSubscriptionRepo() ISubscriptionRepository {
	return &subscriptionRepo{
		db: database.DB(),
	}
}

func (r *subscriptionRepo) Create(sub *model.Subscription) error {
	return r.db.Create(sub).Error
}

func (r *subscriptionRepo) Update(sub *model.Subscription) error {
	return r.db.Save(sub).Error
}
func (r *subscriptionRepo) Delete(userId string, subscriptionId string) error {
	var sub model.Subscription
	if err := r.db.Where("user_id = ? AND subscription_id = ?", userId, subscriptionId).Delete(&sub).Error; err != nil {
		return err
	}
	return nil
}
func (r *subscriptionRepo) FindAllByUserId(userId string) (*[]model.Subscription, error) {
	var sub []model.Subscription
	if err := r.db.Where("user_id = ?", userId).Find(&sub).Error; err != nil {
		return nil, err
	}

	return &sub, nil
}

func (r *subscriptionRepo) FindByUserIdAndSubId(userId string, subscriptionId string) (*model.Subscription, error) {
	var sub model.Subscription
	if err := r.db.Where("user_id = ? AND subscription_id = ?", userId, subscriptionId).First(&sub).Error; err != nil {
		return nil, err
	}

	return &sub, nil
}
