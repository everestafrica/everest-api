package services

import (
	"errors"
	"github.com/everestafrica/everest-api/internal/commons/types"
	"github.com/everestafrica/everest-api/internal/external/crypto"
	"github.com/everestafrica/everest-api/internal/models"
	"github.com/everestafrica/everest-api/internal/repositories"
	"time"
)

type ICryptoService interface {
	SetWallet(coin types.CryptoSymbol, address, userId string) error
	DeleteWallet(coin types.CryptoSymbol, userId string) error
}

type cryptoService struct {
	userRepo          repositories.IUserRepository
	cryptoDetailsRepo repositories.ICryptoDetailsRepository
	cryptoTrxRepo     repositories.ICryptoTransactionRepository
}

// NewCryptoService will instantiate CryptoService
func NewCryptoService() ICryptoService {
	return &cryptoService{
		userRepo:          repositories.NewUserRepo(),
		cryptoDetailsRepo: repositories.NewCryptoDetailsRepo(),
		cryptoTrxRepo:     repositories.NewCryptoTransactionRepo(),
	}
}

func (cs cryptoService) AddWallet(coin types.CryptoSymbol, address string, userId string) error {
	balance, err := crypto.GetBalance(address, coin)
	if err != nil {
		return err
	}

	c := &models.CryptoDetail{
		UserId:        userId,
		WalletAddress: address,
		Balance:       balance.Value,
		Name:          types.CryptoName(GetCoinName(coin)),
		Symbol:        coin,
	}
	err = cs.cryptoDetailsRepo.Create(c)
	if err != nil {
		return err
	}

	transactions, err := crypto.GetTransaction(address, coin)
	if err != nil {
		return err
	}
	for _, transaction := range *transactions {
		trx := &models.CryptoTransaction{
			UserId:        userId,
			WalletAddress: address,
			Name:          types.CryptoName(GetCoinName(coin)),
			Symbol:        coin,
			Value:         transaction.Value,
			Date:          transaction.Date,
			Type:          transaction.Type,
		}
		err := cs.cryptoTrxRepo.Create(trx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (cs cryptoService) SetWallet(coin types.CryptoSymbol, address string, userId string) error {
	balance, err := crypto.GetBalance(address, coin)
	if err != nil {
		return err
	}
	c, err := cs.cryptoDetailsRepo.FindByAddressAndSymbol(address, coin)
	if err != nil {
		return err
	}
	c.Balance = balance.Value

	err = cs.cryptoDetailsRepo.Update(c)
	if err != nil {
		return err
	}

	transactions, err := crypto.GetTransaction(address, coin)
	if err != nil {
		return err
	}

	refreshTimeLimit := time.Now().Add(-12 * time.Hour)

	for _, transaction := range *transactions {
		if transaction.Date.Before(refreshTimeLimit) {
			return errors.New("stale transaction")
		}
		trx := &models.CryptoTransaction{
			UserId:        userId,
			WalletAddress: address,
			Name:          types.CryptoName(GetCoinName(coin)),
			Symbol:        coin,
			Value:         transaction.Value,
			Date:          transaction.Date,
			Type:          transaction.Type,
		}
		err := cs.cryptoTrxRepo.Create(trx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (cs cryptoService) DeleteWallet(coin types.CryptoSymbol, userId string) error {
	return nil
}

func GetCoinName(coin types.CryptoSymbol) string {
	coins := map[types.CryptoSymbol]string{
		"BTC":  "Bitcoin",
		"ETH":  "Ethereum",
		"BSC":  "Binance Coin",
		"USDT": "Tether",
		"SOL":  "Solana",
		"DOGE": "Dogecoin",
	}
	return coins[coin]
}

// Time between 12 hours ago and now

func CalcTime() {

}
