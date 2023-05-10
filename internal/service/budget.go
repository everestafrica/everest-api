package service

import (
	"errors"
	"github.com/everestafrica/everest-api/internal/commons/types"
	"github.com/everestafrica/everest-api/internal/commons/util"
	"github.com/everestafrica/everest-api/internal/model"
	"github.com/everestafrica/everest-api/internal/repository"
)

type IBudgetService interface {
	GetBudget(month string, year int, userId string) (*model.Budget, error)
	CreateBudget(request *types.CreateBudgetRequest, userId string) error
	UpdateBudget(request *types.UpdateBudgetRequest, userId string) error
	DeleteBudget(month string, year int, userId string) error
	GetAllCategories() (*[]model.Category, error)
}

type budgetService struct {
	userRepo   repository.IUserRepository
	budgetRepo repository.IBudgetRepository
}

// NewBudgetService will instantiate BudgetService
func NewBudgetService() IBudgetService {
	return &budgetService{
		userRepo:   repository.NewUserRepo(),
		budgetRepo: repository.NewBudgetRepo(),
	}
}

func (bs budgetService) GetBudget(month string, year int, userId string) (*model.Budget, error) {
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
	if b != nil {
		return errors.New("budget exists for selected time period")
	}
	for _, v := range request.Categories {
		budget := model.Budget{
			UserId:   userId,
			Category: v.Name,
			Amount:   v.Amount,
			Month:    request.Month,
			Year:     request.Year,
			BudgetId: util.GetUUID(),
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
		budget := model.Budget{
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

func (bs budgetService) GetAllCategories() (*[]model.Category, error) {
	categories, err := bs.budgetRepo.FindAllCategories()
	if err != nil {
		return nil, err
	}
	return categories, nil
}
