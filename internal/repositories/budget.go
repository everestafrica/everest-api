package repositories

import (
	"github.com/everestafrica/everest-api/internal/database"
	"github.com/everestafrica/everest-api/internal/models"
	"gorm.io/gorm"
)

type IBudgetRepository interface {
	Create(budget *models.Budget) error
	Update(budget *models.Budget) error
	FindAllByUserId(userId string) (*[]models.Budget, error)
	FindAllByBudgetId(userId string, budgetId string) (*[]models.Budget, error)
	FindByPeriod(userId string, month string, year int) (*[]models.Budget, error)
	Delete(userId string, month string, year int) error
	CreateCustomCategory(category *models.CustomCategory) error
	DeleteCustomCategory(id string) error
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

func (r *budgetRepo) Create(budget *models.Budget) error {
	return r.db.Create(budget).Error
}

func (r *budgetRepo) Update(budget *models.Budget) error {
	return r.db.Save(budget).Error
}

func (r *budgetRepo) FindAllByUserId(userId string) (*[]models.Budget, error) {
	var budget []models.Budget
	if err := r.db.Where("user_id = ?", userId).Find(&budget).Error; err != nil {
		return nil, err
	}
	return &budget, nil
}

func (r *budgetRepo) FindAllByBudgetId(userId string, budgetId string) (*[]models.Budget, error) {
	var budget []models.Budget
	if err := r.db.Where("user_id = ? AND budget_id = ?", userId, budgetId).First(&budget).Error; err != nil {
		return nil, err
	}

	return &budget, nil
}

func (r *budgetRepo) FindByPeriod(userId, month string, year int) (*[]models.Budget, error) {
	var budget []models.Budget
	if err := r.db.Where("user_id = ? AND month =  ? AND year = ?", userId, month, year).Find(&budget).Error; err != nil {
		return nil, err
	}
	return &budget, nil
}

func (r *budgetRepo) Delete(userId, month string, year int) error {
	var budget models.Budget
	if err := r.db.Where("user_id = ? AND month =  ? AND year = ?", userId, month, year).Delete(&budget).Error; err != nil {
		return err
	}
	return nil
}

func (r *budgetRepo) CreateCustomCategory(category *models.CustomCategory) error {
	return r.db.Create(&category).Error
}

func (r *budgetRepo) DeleteCustomCategory(id string) error {
	var category models.CustomCategory
	if err := r.db.Where("id = ? ", id).Delete(&category).Error; err != nil {
		return err
	}
	return nil
}
