package model

import "github.com/everestafrica/everest-api/internal/commons/types"

type CoinRate struct {
	Name     string
	Symbol   types.CryptoSymbol
	Value    float64
	Currency types.CurrencySymbol
}
