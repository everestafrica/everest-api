package services

import (
	"github.com/everestafrica/everest-api/internal/commons/types"
	"github.com/everestafrica/everest-api/internal/models"
	"github.com/everestafrica/everest-api/internal/repositories"
)

type IDebtService interface {
	AddDebt(request *types.CreateDebtRequest, userId string) error
	GetAllDebts(userId string) (*[]models.Debt, error)
	GetDebtByType(userId string, debtType types.DebtType) (*[]models.Debt, error)
	DeleteDebt(debtId, userId string) error
}

type debtService struct {
	userRepo repositories.IUserRepository
	debtRepo repositories.IDebtRepository
}

func NewDebtService() IDebtService {
	return &debtService{
		userRepo: repositories.NewUserRepo(),
		debtRepo: repositories.NewDebtRepo(),
	}
}

func (ds debtService) AddDebt(request *types.CreateDebtRequest, userId string) error {
	debt := models.Debt{
		UserId:           userId,
		Reason:           request.Reason,
		CounterpartyName: request.CounterpartyName,
		Amount:           request.Amount,
	}
	err := ds.debtRepo.Create(&debt)
	if err != nil {
		return err
	}
	return nil
}

func (ds debtService) GetAllDebts(userId string) (*[]models.Debt, error) {
	debts, err := ds.debtRepo.FindAllByUserId(userId)
	if err != nil {
		return nil, err
	}
	return debts, nil
}

func (ds debtService) GetDebtByType(userId string, debtType types.DebtType) (*[]models.Debt, error) {
	debts, err := ds.debtRepo.FindAllByType(userId, debtType)
	if err != nil {
		return nil, err
	}
	return debts, nil
}

func (ds debtService) DeleteDebt(debtId, userId string) error {
	err := ds.debtRepo.Delete(userId, debtId)
	if err != nil {
		return err
	}
	return nil
}
