package services

import (
	"github.com/everestafrica/everest-api/internal/commons/types"
	"github.com/everestafrica/everest-api/internal/commons/utils"
	"github.com/everestafrica/everest-api/internal/models"
	"github.com/everestafrica/everest-api/internal/repositories"
)

type IDebtService interface {
	AddDebt(request *types.CreateDebtRequest, userId string) error
	UpdateDebt(request *types.UpdateDebtRequest, userId string, debtId string) error
	GetAllDebts(userId string) (*[]models.Debt, error)
	GetDebt(userId string, debtId string) (*models.Debt, error)
	GetDebtsByType(userId string, debtType types.DebtType) (*[]models.Debt, error)
	DeleteDebt(userId string, debtId string) error
}

type debtService struct {
	userRepo repositories.IUserRepository
	debtRepo repositories.IDebtRepository
}

func NewDebtService() IDebtService {
	return &debtService{
		debtRepo: repositories.NewDebtRepo(),
	}
}

func (ds debtService) AddDebt(request *types.CreateDebtRequest, userId string) error {
	debt := models.Debt{
		UserId:           userId,
		Reason:           request.Reason,
		CounterpartyName: request.CounterpartyName,
		Amount:           request.Amount,
		DebtId:           utils.GetUUID(),
	}
	err := ds.debtRepo.Create(&debt)
	if err != nil {
		return err
	}
	return nil
}

func (ds debtService) UpdateDebt(request *types.UpdateDebtRequest, userId string, debtId string) error {
	debt, err := ds.debtRepo.FindByUserIdAndDebtId(userId, debtId)
	if err != nil {
		return err
	}
	if &debt.Due != nil {
		debt.Due = *request.Due
	}
	if &debt.Amount != nil {
		debt.Amount = *request.Amount
	}
	if &debt.CounterpartyName != nil {
		debt.CounterpartyName = *request.CounterpartyName
	}
	if &debt.Reason != nil {
		debt.Reason = *request.Reason
	}
	err = ds.debtRepo.Update(debt)
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

func (ds debtService) GetDebt(userId string, debtId string) (*models.Debt, error) {
	debt, err := ds.debtRepo.FindByUserIdAndDebtId(userId, debtId)
	if err != nil {
		return nil, err
	}
	return debt, nil
}

func (ds debtService) GetDebtsByType(userId string, debtType types.DebtType) (*[]models.Debt, error) {
	debts, err := ds.debtRepo.FindAllByType(userId, debtType)
	if err != nil {
		return nil, err
	}
	return debts, nil
}

func (ds debtService) DeleteDebt(userId string, debtId string) error {
	err := ds.debtRepo.Delete(userId, debtId)
	if err != nil {
		return err
	}
	return nil
}
