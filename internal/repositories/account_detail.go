package repositories

import (
	"github.com/everestafrica/everest-api/internal/database"
	"github.com/everestafrica/everest-api/internal/models"
	"gorm.io/gorm"
)

type IAccountDetailsRepository interface {
	Create(account *models.AccountDetail) error
	Update(account *models.AccountDetail) error
	FindByUserId(accountId string, userId string) (*models.AccountDetail, error)
	ExistsByUserId(userId string) bool
	FindAllByUserId(userId string) (*[]models.AccountDetail, error)
	Delete(accountId string, userId string) error
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

func (r *accountDetailsRepo) FindByUserId(accountId string, userId string) (*models.AccountDetail, error) {
	var account models.AccountDetail
	if err := r.db.Where("user_id = ? AND account_id = ?", userId, accountId).First(&account).Error; err != nil {
		return nil, err
	}

	return &account, nil
}
func (r *accountDetailsRepo) ExistsByUserId(userId string) bool {
	return true
	//var user models.User
	//
	//if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
	//	if err == gorm.ErrRecordNotFound {
	//		return false, nil
	//	}
	//	return false, err
	//}
	//
	//return true, nil
}
func (r *accountDetailsRepo) FindAllByUserId(userId string) (*[]models.AccountDetail, error) {
	var account []models.AccountDetail
	if err := r.db.Where("user_id = ?", userId).Find(&account).Error; err != nil {
		return nil, err
	}

	return &account, nil
}
func (r *accountDetailsRepo) Delete(accountId string, userId string) error {
	var account models.AccountDetail
	if err := r.db.Where("user_id = ? AND account_id =  ?", userId, accountId).Delete(&account).Error; err != nil {
		return err
	}
	return nil
}
