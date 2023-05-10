package repository

import (
	"github.com/everestafrica/everest-api/internal/database"
	"github.com/everestafrica/everest-api/internal/model"
	"gorm.io/gorm"
)

type IMonoUserRepository interface {
	Create(user *model.MonoUser) error
	FindByUserId(userId string) (*model.MonoUser, error)
	FindByMonoId(monoId string) (*model.MonoUser, error)
	Update(user *model.MonoUser) error
	Delete(userId string) error
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

func (r *monoUserRepo) Create(user *model.MonoUser) error {
	return r.db.Create(user).Error
}

func (r *monoUserRepo) FindByUserId(userId string) (*model.MonoUser, error) {
	var user model.MonoUser
	if err := r.db.Where("user_id = ?", userId).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *monoUserRepo) FindByMonoId(userId string) (*model.MonoUser, error) {
	var user model.MonoUser
	if err := r.db.Where("mono_id = ?", userId).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *monoUserRepo) Update(user *model.MonoUser) error {
	return r.db.Save(user).Error
}

func (r *monoUserRepo) Delete(userId string) error {
	var user model.MonoUser
	if err := r.db.Where("user_id = ?", userId).Delete(&user).Error; err != nil {
		return err
	}
	return nil
}
