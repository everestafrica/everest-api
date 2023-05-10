package repository

import (
	"github.com/everestafrica/everest-api/internal/database"
	"github.com/everestafrica/everest-api/internal/model"
	"gorm.io/gorm"
)

type IBudgetRepository interface {
	Create(budget *model.Budget) error
	Update(budget *model.Budget) error
	FindAllCategories() (*[]model.Category, error)
	FindAllByUserId(userId string) (*[]model.Budget, error)
	FindAllByBudgetId(userId string, budgetId string) (*[]model.Budget, error)
	FindByPeriod(userId string, month string, year int) (*model.Budget, error)
	Delete(userId string, month string, year int) error
}

type budgetRepo struct {
	db *gorm.DB
}

// NewBudgetRepo will instantiate Budget Repository
func NewBudgetRepo() IBudgetRepository {
	return &budgetRepo{
		db: database.DB(),
	}
}

func (r *budgetRepo) Create(budget *model.Budget) error {
	return r.db.Create(budget).Error
}

func (r *budgetRepo) Update(budget *model.Budget) error {
	return r.db.Save(budget).Error
}

func (r *budgetRepo) FindAllCategories() (*[]model.Category, error) {
	var categories []model.Category
	if err := r.db.Find(&categories).Error; err != nil {
		return nil, err
	}
	return &categories, nil
}

func (r *budgetRepo) FindAllByUserId(userId string) (*[]model.Budget, error) {
	var budget []model.Budget
	if err := r.db.Where("user_id = ?", userId).Find(&budget).Error; err != nil {
		return nil, err
	}
	return &budget, nil
}

func (r *budgetRepo) FindAllByBudgetId(userId string, budgetId string) (*[]model.Budget, error) {
	var budget []model.Budget
	if err := r.db.Where("user_id = ? AND budget_id = ?", userId, budgetId).First(&budget).Error; err != nil {
		return nil, err
	}

	return &budget, nil
}

func (r *budgetRepo) FindByPeriod(userId, month string, year int) (*model.Budget, error) {
	var budget model.Budget
	if err := r.db.Where("user_id = ? AND month =  ? AND year = ?", userId, month, year).Find(&budget).Error; err != nil {
		return nil, err
	}
	return &budget, nil
}

func (r *budgetRepo) Delete(userId, month string, year int) error {
	var budget model.Budget
	if err := r.db.Where("user_id = ? AND month =  ? AND year = ?", userId, month, year).Delete(&budget).Error; err != nil {
		return err
	}
	return nil
}
