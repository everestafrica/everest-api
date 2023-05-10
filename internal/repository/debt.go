package repository

import (
	"github.com/everestafrica/everest-api/internal/commons/types"
	"github.com/everestafrica/everest-api/internal/database"
	"github.com/everestafrica/everest-api/internal/model"
	"gorm.io/gorm"
)

type IDebtRepository interface {
	Create(debt *model.Debt) error
	Update(debt *model.Debt) error
	Delete(userId, debtId string) error
	FindAllByUserId(userId string) (*[]model.Debt, error)
	FindAllByType(userId string, debtType types.DebtType) (*[]model.Debt, error)
	FindByUserIdAndDebtId(userId string, debtId string) (*model.Debt, error)
}

type debtRepo struct {
	db *gorm.DB
}

// NewDebtRepo will instantiate Debt Repository
func NewDebtRepo() IDebtRepository {
	return &debtRepo{
		db: database.DB(),
	}
}

func (r *debtRepo) Create(debt *model.Debt) error {
	return r.db.Create(debt).Error
}

func (r *debtRepo) Update(debt *model.Debt) error {
	return r.db.Save(debt).Error
}
func (r *debtRepo) Delete(userId, debtId string) error {
	var debt model.Debt
	if err := r.db.Where("user_id = ? AND debt_id = ?", userId, debtId).Delete(&debt).Error; err != nil {
		return err
	}
	return nil
}
func (r *debtRepo) FindAllByUserId(userId string) (*[]model.Debt, error) {
	var debt []model.Debt
	if err := r.db.Where("user_id = ?", userId).Find(&debt).Error; err != nil {
		return nil, err
	}

	return &debt, nil
}

func (r *debtRepo) FindAllByType(userId string, debtType types.DebtType) (*[]model.Debt, error) {
	var debt []model.Debt
	if err := r.db.Where("user_id = ? AND type = ?", userId, debtType).Find(&debt).Error; err != nil {
		return nil, err
	}

	return &debt, nil
}

func (r *debtRepo) FindByUserIdAndDebtId(userId string, debtId string) (*model.Debt, error) {
	var debt model.Debt
	if err := r.db.Where("user_id = ? AND debt_id = ?", userId, debtId).First(&debt).Error; err != nil {
		return nil, err
	}

	return &debt, nil
}
