package repositories

import (
	"github.com/everestafrica/everest-api/internal/database"
	"github.com/everestafrica/everest-api/internal/models"
	"gorm.io/gorm"
)

type IAccountDetailsRepository interface {
	Create(account *models.AccountDetail) error
	Update(account *models.AccountDetail) error
	FindByUserId(userId string) (*models.AccountDetail, error)
}

type accountDetailsRepo struct {
	db *gorm.DB
}

// NewAccountDetailsRepo will instantiate AccountDetails Repository
func NewAccountDetailsRepo() IAccountDetailsRepository {
	return &accountDetailsRepo{
		db: database.DB(),
	}
}

func (r *accountDetailsRepo) Create(account *models.AccountDetail) error {
	return r.db.Create(account).Error
}

func (r *accountDetailsRepo) Update(account *models.AccountDetail) error {
	return r.db.Save(account).Error
}

func (r *accountDetailsRepo) FindByUserId(userId string) (*models.AccountDetail, error) {
	var account models.AccountDetail
	if err := r.db.Where("user_id = ?", userId).First(&account).Error; err != nil {
		return nil, err
	}

	return &account, nil
}
