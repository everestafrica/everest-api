package models

import "github.com/everestafrica/everest-api/internal/commons/types"

type CryptoDetail struct {
	Base
	UserId        string             `json:"user_id"`
	WalletAddress string             `json:"wallet_address"`
	Balance       float64            `json:"balance"`
	Name          types.CryptoName   `json:"name"`
	Symbol        types.CryptoSymbol `json:"symbol"`
}

type CryptoTransaction struct {
	Base
	UserId        string                `json:"user_id"`
	WalletAddress string                `json:"wallet_address"`
	Hash          string                `json:"hash"`
	Name          types.CryptoName      `json:"name"`
	Symbol        types.CryptoSymbol    `json:"symbol"`
	Fees          string                `json:"fees"`
	Value         string                `json:"value"`
	Date          string                `json:"date"`
	Type          types.TransactionType `json:"type"`
}
