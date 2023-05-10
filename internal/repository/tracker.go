package repository

import (
	"github.com/everestafrica/everest-api/internal/database"
	"github.com/everestafrica/everest-api/internal/model"
	"gorm.io/gorm"
)

type ITrackerRepository interface {
	Create(tracker *model.Tracker) error
	Update(tracker *model.Tracker) error
	FindByUserId(userId string) (*model.Tracker, error)
	//Delete(userId, trackerId string) error
}

type trackerRepo struct {
	db *gorm.DB
}

// NewTrackerRepo will instantiate Tracker Repository
func NewTrackerRepo() ITrackerRepository {
	return &trackerRepo{
		db: database.DB(),
	}
}

func (r trackerRepo) Create(tracker *model.Tracker) error {
	return r.db.Create(tracker).Error
}

func (r trackerRepo) Update(tracker *model.Tracker) error {
	return r.db.Save(tracker).Error
}

func (r trackerRepo) FindByUserId(userId string) (*model.Tracker, error) {
	var tracker model.Tracker
	if err := r.db.Where("user_id = ?", userId).First(&tracker).Error; err != nil {
		return nil, err
	}
	return &tracker, nil
}
