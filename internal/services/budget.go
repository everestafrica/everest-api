package services

import (
	"errors"
	"github.com/everestafrica/everest-api/internal/commons/types"
	"github.com/everestafrica/everest-api/internal/models"
	"github.com/everestafrica/everest-api/internal/repositories"
)

type IBudgetService interface {
	CreateBudget(request *types.CreateBudgetRequest, userId string) error
	UpdateBudget(userId string, category int, request *types.UpdateBudgetRequest) error
	DeleteBudget(budgetId string, userId string) error
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

func (bs budgetService) CreateBudget(request *types.CreateBudgetRequest, userId string) error {
	tracker, err := bs.trackerRepo.FindByUserId(userId)
	if tracker.CreatedBudget {
		return errors.New("budget created already")
	}
	budget := models.Budget{
		//BudgetId:    uuid.NewString(),
		UserId:      userId,
		TotalAmount: request.TotalAmount,
		//StartDay:    request.Start,
		//EndDay:      request.End,
	}
	err = bs.budgetRepo.Create(&budget)
	if err != nil {
		return err
	}
	for _, v := range request.Categories {
		category := models.Category{
			BudgetId: budget.ID,
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

func (bs budgetService) UpdateBudget(userId string, category int, request *types.UpdateBudgetRequest) error {
	//cat, err := bs.budgetRepo.FindByUserIdAndCategoryId(userId, category)
	//if err != nil {
	//	return err
	//}
	//for _, v := range request.Categories {
	//	category := models.Category{
	//		BudgetId: budget.BudgetId,
	//		Name:     v.Name,
	//		Amount:   v.Amount,
	//	}
	//	err := bs.budgetRepo.CreateCategory(&category)
	//	if err != nil {
	//		return err
	//	}
	//}

	return nil
}

func (bs budgetService) DeleteBudget(budgetId string, userId string) error {
	err := bs.budgetRepo.Delete(userId, budgetId)
	if err != nil {
		return err
	}
	return nil
}
