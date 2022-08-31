package services

import (
	"errors"
	"github.com/everestafrica/everest-api/internal/commons/types"
	"github.com/everestafrica/everest-api/internal/models"
	"github.com/everestafrica/everest-api/internal/repositories"
)

type IBudgetService interface {
	GetBudget(month string, year int, userId string) (*[]models.Budget, error)
	CreateBudget(request *types.CreateBudgetRequest, userId string) error
	UpdateBudget(request *types.UpdateBudgetRequest, userId string) error
	DeleteBudget(month string, year int, userId string) error
}

type budgetService struct {
	userRepo    repositories.IUserRepository
	trackerRepo repositories.ITrackerRepository
	budgetRepo  repositories.IBudgetRepository
}

// NewBudgetService will instantiate BudgetService
func NewBudgetService() IBudgetService {
	return &budgetService{
		userRepo:    repositories.NewUserRepo(),
		trackerRepo: repositories.NewTrackerRepo(),
		budgetRepo:  repositories.NewBudgetRepo(),
	}
}

func (bs budgetService) GetBudget(month string, year int, userId string) (*[]models.Budget, error) {
	budget, err := bs.budgetRepo.FindByPeriod(userId, month, year)
	if err != nil {
		return nil, err
	}
	return budget, nil
}

func (bs budgetService) CreateBudget(request *types.CreateBudgetRequest, userId string) error {
	b, err := bs.budgetRepo.FindByPeriod(userId, request.Month, request.Year)
	if err != nil {
		return err
	}
	if len(*b) > 0 {
		return errors.New("budget exists for selected time period")
	}
	for _, v := range request.Categories {
		budget := models.Budget{
			UserId:   userId,
			Category: v.Name,
			Amount:   v.Amount,
			Month:    request.Month,
			Year:     request.Year,
		}
		err = bs.budgetRepo.Create(&budget)
		if err != nil {
			return err
		}
	}

	return nil
}

func (bs budgetService) UpdateBudget(request *types.UpdateBudgetRequest, userId string) error {
	err := bs.budgetRepo.Delete(userId, request.Month, request.Year)
	if err != nil {
		return err
	}

	for _, v := range request.Categories {
		budget := models.Budget{
			UserId:   userId,
			Category: v.Name,
			Amount:   v.Amount,
			Month:    request.Month,
			Year:     request.Year,
		}
		err = bs.budgetRepo.Create(&budget)
		if err != nil {
			return err
		}
	}
	return nil
}

func (bs budgetService) DeleteBudget(month string, year int, userId string) error {
	err := bs.budgetRepo.Delete(userId, month, year)
	if err != nil {
		return err
	}
	return nil
}
