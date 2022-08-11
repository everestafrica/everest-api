package repositories

import (
	"github.com/everestafrica/everest-api/internal/database"
	"github.com/everestafrica/everest-api/internal/models"
	"gorm.io/gorm"
)

type ISubscriptionRepository interface {
	Create(sub *models.Subscription) error
	Update(sub *models.Subscription) error
	Delete(userId, subscriptionId string) error
	FindAllByUserId(userId string) (*[]models.Subscription, error)
	FindByUserIdAndSubId(userId string, subscriptionId int) (*models.Subscription, error)
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

func (r *subscriptionRepo) Create(sub *models.Subscription) error {
	return r.db.Create(sub).Error
}

func (r *subscriptionRepo) Update(sub *models.Subscription) error {
	return r.db.Save(sub).Error
}
func (r *subscriptionRepo) Delete(userId, subscriptionId string) error {
	var sub models.Subscription
	if err := r.db.Where("user_id = ? AND id =  ?", userId, subscriptionId).Delete(&sub).Error; err != nil {
		return err
	}
	return nil
}
func (r *subscriptionRepo) FindAllByUserId(userId string) (*[]models.Subscription, error) {
	var sub []models.Subscription
	if err := r.db.Where("user_id = ?", userId).Find(&sub).Error; err != nil {
		return nil, err
	}

	return &sub, nil
}

func (r *subscriptionRepo) FindByUserIdAndSubId(userId string, subscriptionId int) (*models.Subscription, error) {
	var sub models.Subscription
	if err := r.db.Where("user_id = ? AND id = ?", userId, subscriptionId).First(&sub).Error; err != nil {
		return nil, err
	}

	return &sub, nil
}
