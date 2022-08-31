package services

import (
	"github.com/everestafrica/everest-api/internal/commons/types"
	"github.com/everestafrica/everest-api/internal/external/crypto"
	"github.com/everestafrica/everest-api/internal/models"
	"github.com/everestafrica/everest-api/internal/repositories"
	"strconv"
)

type ICryptoService interface {
	AddWallet(coin types.CryptoSymbol, address, userId string) error
}

type cryptoService struct {
	userRepo   repositories.IUserRepository
	cryptoRepo repositories.ICryptoDetailsRepository
}

// NewCryptoService will instantiate CryptoService
func NewCryptoService() ICryptoService {
	return &cryptoService{
		userRepo:   repositories.NewUserRepo(),
		cryptoRepo: repositories.NewCryptoDetailsRepo(),
	}
}

func (cs cryptoService) AddWallet(coin types.CryptoSymbol, address string, userId string) error {
	balance, err := crypto.GetBalance(address, coin)
	if err != nil {
		return err
	}
	bal, err := strconv.Atoi(*balance)
	if err != err {
		return err
	}
	c := &models.CryptoDetail{
		UserId:        userId,
		WalletAddress: address,
		Balance:       bal,
		Name:          types.CryptoName(GetCoinName(coin)),
		Symbol:        coin,
	}
	err = cs.cryptoRepo.Create(c)
	if err != nil {
		return err
	}
	return nil
}

func (cs cryptoService) DeleteWallet(coin types.CryptoSymbol, userId string) error {
	return nil
}

func GetCoinName(coin types.CryptoSymbol) string {
	return ""
}
