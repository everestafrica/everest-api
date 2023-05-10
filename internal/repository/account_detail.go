package repository

import (
	"github.com/everestafrica/everest-api/internal/database"
	"github.com/everestafrica/everest-api/internal/model"
	"gorm.io/gorm"
)

type IAccountDetailsRepository interface {
	Create(account *model.AccountDetail) error
	Update(account *model.AccountDetail) error
	FindByAccountId(accountId string) (*model.AccountDetail, error)
	ExistsByAccountInstitution(institution string, userId string) bool
	FindAllByUserId(userId string) (*[]model.AccountDetail, error)
	Delete(accountId string) error
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

func (r *accountDetailsRepo) Create(account *model.AccountDetail) error {
	return r.db.Create(account).Error
}

func (r *accountDetailsRepo) Update(account *model.AccountDetail) error {
	return r.db.Save(account).Error
}

func (r *accountDetailsRepo) FindByAccountId(accountId string) (*model.AccountDetail, error) {
	var account model.AccountDetail
	if err := r.db.Where("account_id = ?", accountId).First(&account).Error; err != nil {
		return nil, err
	}
	return &account, nil
}

func (r *accountDetailsRepo) ExistsByAccountInstitution(institution string, userId string) bool {
	var account model.AccountDetail

	if err := r.db.Where("user_id = ? AND institution = ?", userId, institution).First(&account).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false
		}
		return false
	}

	return true
}
func (r *accountDetailsRepo) FindAllByUserId(userId string) (*[]model.AccountDetail, error) {
	var account []model.AccountDetail
	if err := r.db.Where("user_id = ?", userId).Find(&account).Error; err != nil {
		return nil, err
	}

	return &account, nil
}
func (r *accountDetailsRepo) Delete(accountId string) error {
	var account model.AccountDetail
	if err := r.db.Where("account_id =  ?", accountId).Delete(&account).Error; err != nil {
		return err
	}
	return nil
}
