package services

import (
	"errors"
	"fmt"
	"github.com/everestafrica/everest-api/internal/commons/log"
	"github.com/everestafrica/everest-api/internal/commons/types"
	"github.com/everestafrica/everest-api/internal/commons/utils"
	"github.com/everestafrica/everest-api/internal/external/crypto"
	"github.com/everestafrica/everest-api/internal/models"
	"github.com/everestafrica/everest-api/internal/repositories"
	"time"
)

//0x639dAec464Bfd4fc5173FD957d6f4363cC96282f
//bc1q673dyteuz89en8lhhwpjvfdyan78ragj42034n
//D2aSGBsGESJeyeFuDUd8dpaY2PEeu97kgozPaFGa3g7u

type ICoinService interface {
	GetAllCoinWallets(userId string) (*[]models.CoinWallet, error)
	GetCoinWallet(id string) (*models.CoinWallet, error)
	AddCoinWallet(symbol types.CoinSymbol, address string, userId string) error
	UpdateCoinWallet(symbol types.CoinSymbol, address, userId string) error
	DeleteCoinWallet(symbol types.CoinSymbol, address string, userId string) error
	GetAllTransactions(userId string) (*[]models.CoinTransaction, error)
	GetTransaction(hash string) (*models.CoinTransaction, error)

	GetInflow(dateRange types.DateRange, userId string) (*types.TxnFlowResponse, error)
	GetOutflow(dateRange types.DateRange, userId string) (*types.TxnFlowResponse, error)
}

type coinService struct {
	userRepo     repositories.IUserRepository
	coinRepo     repositories.ICoinRepository
	coinTrxRepo  repositories.ICoinTransactionRepository
	assetService assetService
}

// NewCoinService will instantiate CoinService
func NewCoinService() ICoinService {
	return &coinService{
		userRepo:    repositories.NewUserRepo(),
		coinRepo:    repositories.NewCoinRepo(),
		coinTrxRepo: repositories.NewCryptoTransactionRepo(),
	}
}

func (cs coinService) GetAllCoinWallets(userId string) (*[]models.CoinWallet, error) {
	coins, err := cs.coinRepo.FindByUserId(userId)
	if err != nil {
		return nil, err
	}
	return coins, nil
}

func (cs coinService) GetCoinWallet(id string) (*models.CoinWallet, error) {
	wallet, err := cs.coinRepo.FindById(id)
	if err != nil {
		return nil, err
	}
	return wallet, nil
}

func (cs coinService) AddCoinWallet(symbol types.CoinSymbol, address string, userId string) error {
	balance, err := crypto.GetBalance(address, symbol)
	if err != nil {
		log.Errorf("error getting balance for %s", address)
		return err
	}

	c := &models.CoinWallet{
		UserId:        userId,
		WalletAddress: address,
		Balance:       balance.Value,
		Name:          types.CoinName(utils.GetCoinName(symbol)),
		Symbol:        symbol,
	}
	err = cs.coinRepo.Create(c)
	if err != nil {
		log.Errorf("error creating coin wallet for %s", address)
		return err
	}

	//err = cs.assetService.AddAsset(string(symbol), true, userId)
	//if err != nil {
	//	log.Errorf("error adding asset for %s", address)
	//	return err
	//}

	transactions, err := crypto.GetTransaction(address, symbol)
	if err != nil {
		log.Errorf("error getting transactions for %s", address)
		return err
	}
	for _, transaction := range *transactions {
		trx := &models.CoinTransaction{
			UserId:   userId,
			WalletId: c.Id,
			Value:    transaction.Value,
			Date:     transaction.Date,
			Type:     transaction.Type,
		}
		err = cs.coinTrxRepo.Create(trx)
		if err != nil {
			log.Errorf("error creating coin transaction for %s", address)
			return err
		}
	}

	return nil
}

func (cs coinService) UpdateCoinWallet(symbol types.CoinSymbol, address string, userId string) error {
	balance, err := crypto.GetBalance(address, symbol)
	if err != nil {
		return err
	}
	c, err := cs.coinRepo.FindByAddressAndSymbol(address, symbol)
	if err != nil {
		return err
	}
	c.Balance = balance.Value

	err = cs.coinRepo.Update(c)
	if err != nil {
		return err
	}

	transactions, err := crypto.GetTransaction(address, symbol)
	if err != nil {
		return err
	}

	refreshTimeLimit := time.Now().Add(-12 * time.Hour)

	for _, transaction := range *transactions {
		if transaction.Date.Before(refreshTimeLimit) {
			return errors.New("stale transaction")
		}
		trx := &models.CoinTransaction{
			UserId:   userId,
			WalletId: c.Id,
			Value:    transaction.Value,
			Date:     transaction.Date,
			Type:     transaction.Type,
		}
		err = cs.coinTrxRepo.Create(trx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (cs coinService) DeleteCoinWallet(symbol types.CoinSymbol, address string, userId string) error {
	err := cs.coinRepo.Delete(userId, symbol, address)
	if err != nil {
		return err
	}
	return nil
}

func (cs coinService) GetAllTransactions(userId string) (*[]models.CoinTransaction, error) {
	transactions, err := cs.coinTrxRepo.FindByUserId(userId)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (cs coinService) GetTransaction(hash string) (*models.CoinTransaction, error) {
	transactions, err := cs.coinTrxRepo.FindTransaction(hash)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (cs coinService) GetInflow(dateRange types.DateRange, userId string) (*types.TxnFlowResponse, error) {
	transactions, err := cs.coinTrxRepo.FindAllTxnFlow(types.Credit, dateRange, userId)
	if err != nil {
		return nil, err
	}
	var inflow float64
	for _, v := range *transactions {
		inflow += v.Amount
	}
	result := &types.TxnFlowResponse{
		Total:     inflow,
		DateRange: fmt.Sprintf("from: %s - to: %s", dateRange.From, dateRange.To),
	}
	return result, err
}
func (cs coinService) GetOutflow(dateRange types.DateRange, userId string) (*types.TxnFlowResponse, error) {
	transactions, err := cs.coinTrxRepo.FindAllTxnFlow(types.Debit, dateRange, userId)
	if err != nil {
		return nil, err
	}
	var outflow float64
	for _, v := range *transactions {
		outflow += v.Amount

	}
	result := &types.TxnFlowResponse{
		Total:     outflow,
		DateRange: fmt.Sprintf("%s - %s", dateRange.From, dateRange.To),
	}
	return result, err
}
