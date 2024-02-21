package repositories

import (
	"github.com/everestafrica/everest-api/internal/database"
	"github.com/everestafrica/everest-api/internal/models"
	"gorm.io/gorm"
)

type ICashAccountRepository interface {
	Create(cashAccount *models.CashAccount) error
	Update(cashAccount *models.CashAccount) error
	FindByAccountId(cashAccountId string) (*models.CashAccount, error)
	ExistsByAccountInstitution(institution string, userId string) bool
	FindAllByUserId(userId string) (*[]models.CashAccount, error)
	Delete(cashAccountId string) error
}

type cashRepo struct {
	db *gorm.DB
}

// NewCashRepo will instantiate Cash Repository
func NewCashRepo() ICashAccountRepository {
	return &cashRepo{
		db: database.DB(),
	}
}

func (r *cashRepo) Create(cashAccount *models.CashAccount) error {
	return r.db.Create(cashAccount).Error
}

func (r *cashRepo) Update(cashAccount *models.CashAccount) error {
	return r.db.Save(cashAccount).Error
}

func (r *cashRepo) FindByAccountId(cashAccountId string) (*models.CashAccount, error) {
	var cashAccount models.CashAccount
	if err := r.db.Where("account_id = ?", cashAccountId).First(&cashAccount).Error; err != nil {
		return nil, err
	}
	return &cashAccount, nil
}

func (r *cashRepo) ExistsByAccountInstitution(institution string, userId string) bool {
	var cashAccount models.CashAccount

	if err := r.db.Where("user_id = ? AND institution = ?", userId, institution).First(&cashAccount).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false
		}
		return false
	}

	return true
}
func (r *cashRepo) FindAllByUserId(userId string) (*[]models.CashAccount, error) {
	var cashAccount []models.CashAccount
	if err := r.db.Where("user_id = ?", userId).Find(&cashAccount).Error; err != nil {
		return nil, err
	}

	return &cashAccount, nil
}
func (r *cashRepo) Delete(cashAccountId string) error {
	var cashAccount models.CashAccount
	if err := r.db.Where("account_id =  ?", cashAccountId).Delete(&cashAccount).Error; err != nil {
		return err
	}
	return nil
}
