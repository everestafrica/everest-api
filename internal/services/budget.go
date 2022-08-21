package services

import (
	"github.com/everestafrica/everest-api/internal/commons/types"
	"github.com/everestafrica/everest-api/internal/models"
	"github.com/everestafrica/everest-api/internal/repositories"
	"github.com/google/uuid"
)

type IBudgetService interface {
	CreateBudget(request *types.BudgetRequest, userId string) error
	//	UpdateBudget(request interface{}) error
	//	DeleteBudget(budgetId string) error
}

type budgetService struct {
	userRepo   repositories.IUserRepository
	budgetRepo repositories.IBudgetRepository
}

// NewBudgetService will instantiate BudgetService
func NewBudgetService() IBudgetService {
	return &budgetService{
		userRepo:   repositories.NewUserRepo(),
		budgetRepo: repositories.NewBudgetRepo(),
	}
}

func (bs budgetService) CreateBudget(request *types.BudgetRequest, userId string) error {
	budget := models.Budget{
		BudgetId:    uuid.NewString(),
		UserId:      userId,
		TotalAmount: request.TotalAmount,
		StartDay:    request.Start,
		EndDay:      request.End,
	}
	err := bs.budgetRepo.Create(&budget)
	if err != nil {
		return err
	}
	for _, v := range request.Categories {
		category := models.Category{
			BudgetId: budget.BudgetId,
			Name:     v.Name,
			Amount:   v.Amount,
		}
		err := bs.budgetRepo.CreateCategory(&category)
		if err != nil {
			return err
		}
	}

	return nil
}
