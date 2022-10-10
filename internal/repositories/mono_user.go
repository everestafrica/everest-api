package repositories

import (
	"github.com/everestafrica/everest-api/internal/database"
	"github.com/everestafrica/everest-api/internal/models"
	"gorm.io/gorm"
)

type IMonoUserRepository interface {
	Create(user *models.MonoUser) error
	FindByUserId(userId string) (*models.MonoUser, error)
	FindByMonoId(monoId string) (*models.MonoUser, error)
	Update(user *models.MonoUser) error
}

type monoUserRepo struct {
	db *gorm.DB
}

// NewMonoUserRepo will instantiate User Repository
func NewMonoUserRepo() IMonoUserRepository {
	return &monoUserRepo{
		db: database.DB(),
	}
}

func (r *monoUserRepo) Create(user *models.MonoUser) error {
	return r.db.Create(user).Error
}

func (r *monoUserRepo) FindByUserId(userId string) (*models.MonoUser, error) {
	var user models.MonoUser
	if err := r.db.Where("user_id = ?", userId).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *monoUserRepo) FindByMonoId(userId string) (*models.MonoUser, error) {
	var user models.MonoUser
	if err := r.db.Where("mono_id = ?", userId).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *monoUserRepo) Update(user *models.MonoUser) error {
	return r.db.Save(user).Error
}
