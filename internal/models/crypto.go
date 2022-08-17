package models

import "github.com/everestafrica/everest-api/internal/commons/types"

type CryptoDetail struct {
	Base
	UserId        string             `json:"user_id"`
	WalletAddress string             `json:"wallet_address"`
	Balance       int                `json:"balance"`
	Name          types.CryptoName   `json:"name"`
	Symbol        types.CryptoSymbol `json:"symbol"`
}

type CryptoTransaction struct {
	Base
	UserId string             `json:"user_id"`
	Hash   string             `json:"hash"`
	Name   types.CryptoName   `json:"name"`
	Symbol types.CryptoSymbol `json:"symbol"`
	Amount int                `json:"amount"`
	Value  int                `json:"value"`
	Date   string             `json:"date"`
	Type   string             `json:"type"`
}
