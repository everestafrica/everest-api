package models

import (
	"github.com/everestafrica/everest-api/internal/commons/types"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type CoinWallet struct {
	Base
	UserId        string           `json:"-"`
	Id            string           `json:"id"`
	WalletAddress string           `json:"wallet_address"`
	Balance       float64          `json:"balance"`
	Name          types.CoinName   `json:"name"`
	Symbol        types.CoinSymbol `json:"symbol"`
}

type CoinTransaction struct {
	Base
	UserId   string                    `json:"-"`
	Id       string                    `json:"id"`
	WalletId string                    `json:"wallet_id"`
	Hash     string                    `json:"hash"`
	Fees     string                    `json:"fees"`
	Value    float64                   `json:"value"`
	Amount   float64                   `json:"amount"`
	Category types.TransactionCategory `json:"category"`
	Type     types.TransactionType     `json:"type"`
	Date     time.Time                 `json:"date"`
}

func (u *CoinWallet) BeforeCreate(tx *gorm.DB) (err error) {
	u.Id = uuid.NewString()
	return nil
}

func (u *CoinTransaction) BeforeCreate(tx *gorm.DB) (err error) {
	u.Id = uuid.NewString()
	return nil
}
