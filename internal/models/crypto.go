package models

import (
	"github.com/everestafrica/everest-api/internal/commons/types"
	"time"
)

type CryptoDetail struct {
	Base
	UserId        string             `json:"-"`
	WalletAddress string             `json:"wallet_address"`
	Balance       float64            `json:"balance"`
	Name          types.CryptoName   `json:"name"`
	Symbol        types.CryptoSymbol `json:"symbol"`
}

type CryptoTransaction struct {
	Base
	UserId        string                    `json:"-"`
	WalletAddress string                    `json:"wallet_address"`
	Hash          string                    `json:"hash"`
	Name          types.CryptoName          `json:"name"`
	Symbol        types.CryptoSymbol        `json:"symbol"`
	Fees          string                    `json:"fees"`
	Value         float64                   `json:"value"`
	Amount        float64                   `json:"amount"`
	Category      types.TransactionCategory `json:"category"`
	Type          types.TransactionType     `json:"type"`
	Date          time.Time                 `json:"date"`
}
