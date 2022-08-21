package repositories

import (
	"github.com/everestafrica/everest-api/internal/database"
	"github.com/everestafrica/everest-api/internal/models"
	"gorm.io/gorm"
)

type IBudgetRepository interface {
	Create(budget *models.Budget) error
	Update(budget *models.Budget) error
	CreateCategory(category *models.Category) error
	UpdateCategory(category *models.Category) error
	FindAllByUserId(userId string) (*[]models.Budget, error)
	Delete(userId, budgetId string) error
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

func (r budgetRepo) Create(budget *models.Budget) error {
	return r.db.Create(budget).Error
}

func (r budgetRepo) Update(budget *models.Budget) error {
	return r.db.Save(budget).Error
}

func (r budgetRepo) CreateCategory(category *models.Category) error {
	return r.db.Create(category).Error
}

func (r budgetRepo) UpdateCategory(category *models.Category) error {
	return r.db.Save(category).Error
}

func (r budgetRepo) FindAllByUserId(userId string) (*[]models.Budget, error) {
	var budget []models.Budget
	if err := r.db.Where("user_id = ?", userId).First(&budget).Error; err != nil {
		return nil, err
	}
	return &budget, nil
}

func (r *budgetRepo) FindByUserIdAndBudgetId(userId string, budgetId int) (*models.Budget, error) {
	var budget models.Budget
	if err := r.db.Where("user_id = ? AND budget_id = ?", userId, budgetId).First(&budget).Error; err != nil {
		return nil, err
	}

	return &budget, nil
}

func (r *budgetRepo) Delete(userId, budgetId string) error {
	var budget models.Budget
	if err := r.db.Where("user_id = ? AND budget_id =  ?", userId, budgetId).Delete(&budget).Error; err != nil {
		return err
	}
	return nil
}
