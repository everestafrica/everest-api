package service

import (
	"errors"
	"fmt"
	"github.com/everestafrica/everest-api/internal/commons/types"
	"github.com/everestafrica/everest-api/internal/commons/util"
	"github.com/everestafrica/everest-api/internal/external/asset"
	"github.com/everestafrica/everest-api/internal/external/crypto"
	"github.com/everestafrica/everest-api/internal/model"
	"github.com/everestafrica/everest-api/internal/repository"
	"time"
)

type ICryptoService interface {
	GetAllWallets(userId string) (*[]model.CryptoDetail, error)
	GetWallet(id string) (*model.CryptoDetail, error)
	AddWallet(symbol types.CryptoSymbol, address string, userId string) error
	UpdateWallet(symbol types.CryptoSymbol, address, userId string) error
	DeleteWallet(symbol types.CryptoSymbol, address string, userId string) error
	GetAllTransactions(userId string) (*[]model.CryptoTransaction, error)
	GetTransaction(hash string) (*model.CryptoTransaction, error)

	GetInflow(dateRange types.DateRange, userId string) (*types.TxnFlowResponse, error)
	GetOutflow(dateRange types.DateRange, userId string) (*types.TxnFlowResponse, error)
}

type cryptoService struct {
	userRepo          repository.IUserRepository
	cryptoDetailsRepo repository.ICryptoDetailsRepository
	cryptoTrxRepo     repository.ICryptoTransactionRepository
	assetService      assetService
}

// NewCryptoService will instantiate CryptoService
func NewCryptoService() ICryptoService {
	return &cryptoService{
		userRepo:          repository.NewUserRepo(),
		cryptoDetailsRepo: repository.NewCryptoDetailsRepo(),
		cryptoTrxRepo:     repository.NewCryptoTransactionRepo(),
	}
}

func (cs cryptoService) GetAllWallets(userId string) (*[]model.CryptoDetail, error) {
	wallets, err := cs.cryptoDetailsRepo.FindByUserId(userId)
	if err != nil {
		return nil, err
	}
	return wallets, nil
}

func (cs cryptoService) GetWallet(id string) (*model.CryptoDetail, error) {
	wallet, err := cs.cryptoDetailsRepo.FindById(id)
	if err != nil {
		return nil, err
	}
	return wallet, nil
}

func (cs cryptoService) AddWallet(symbol types.CryptoSymbol, address string, userId string) error {
	balance, err := crypto.GetBalance(address, symbol)
	if err != nil {
		return err
	}

	c := &model.CryptoDetail{
		UserId:        userId,
		WalletId:      util.GetUUID(),
		WalletAddress: address,
		Balance:       balance.Value,
		Name:          types.CryptoName(asset.GetCoinName(symbol)),
		Symbol:        symbol,
	}
	err = cs.cryptoDetailsRepo.Create(c)
	if err != nil {
		return err
	}

	err = cs.assetService.AddAsset(string(symbol), true, userId)
	if err != nil {
		return err
	}

	transactions, err := crypto.GetTransaction(address, symbol)
	if err != nil {
		return err
	}
	for _, transaction := range *transactions {
		trx := &model.CryptoTransaction{
			UserId:        userId,
			WalletAddress: address,
			Name:          types.CryptoName(asset.GetCoinName(symbol)),
			Symbol:        symbol,
			Value:         transaction.Value,
			Date:          transaction.Date,
			Type:          transaction.Type,
		}
		err = cs.cryptoTrxRepo.Create(trx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (cs cryptoService) UpdateWallet(symbol types.CryptoSymbol, address string, userId string) error {
	balance, err := crypto.GetBalance(address, symbol)
	if err != nil {
		return err
	}
	c, err := cs.cryptoDetailsRepo.FindByAddressAndSymbol(address, symbol)
	if err != nil {
		return err
	}
	c.Balance = balance.Value

	err = cs.cryptoDetailsRepo.Update(c)
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
		trx := &model.CryptoTransaction{
			UserId:        userId,
			WalletAddress: address,
			Name:          types.CryptoName(asset.GetCoinName(symbol)),
			Symbol:        symbol,
			Value:         transaction.Value,
			Date:          transaction.Date,
			Type:          transaction.Type,
		}
		err = cs.cryptoTrxRepo.Create(trx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (cs cryptoService) DeleteWallet(symbol types.CryptoSymbol, address string, userId string) error {
	err := cs.cryptoDetailsRepo.Delete(userId, symbol, address)
	if err != nil {
		return err
	}
	return nil
}

func (cs cryptoService) GetAllTransactions(userId string) (*[]model.CryptoTransaction, error) {
	transactions, err := cs.cryptoTrxRepo.FindByUserId(userId)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (cs cryptoService) GetTransaction(hash string) (*model.CryptoTransaction, error) {
	transactions, err := cs.cryptoTrxRepo.FindTransaction(hash)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (cs cryptoService) GetInflow(dateRange types.DateRange, userId string) (*types.TxnFlowResponse, error) {
	transactions, err := cs.cryptoTrxRepo.FindAllTxnFlow(types.Credit, dateRange, userId)
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
func (cs cryptoService) GetOutflow(dateRange types.DateRange, userId string) (*types.TxnFlowResponse, error) {
	transactions, err := cs.cryptoTrxRepo.FindAllTxnFlow(types.Debit, dateRange, userId)
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
