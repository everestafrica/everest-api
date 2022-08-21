package repositories

import (
	"github.com/everestafrica/everest-api/internal/database"
	"github.com/everestafrica/everest-api/internal/models"
	"gorm.io/gorm"
)

type ITrackerRepository interface {
	Create(tracker *models.Tracker) error
	Update(tracker *models.Tracker) error
	FindAllByUserId(userId string) (*[]models.Tracker, error)
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

func (r trackerRepo) Create(tracker *models.Tracker) error {
	return r.db.Create(tracker).Error
}

func (r trackerRepo) Update(tracker *models.Tracker) error {
	return r.db.Save(tracker).Error
}

func (r trackerRepo) FindAllByUserId(userId string) (*[]models.Tracker, error) {
	var tracker []models.Tracker
	if err := r.db.Where("user_id = ?", userId).Find(&tracker).Error; err != nil {
		return nil, err
	}
	return &tracker, nil
}
