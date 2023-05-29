package models

import "github.com/everestafrica/everest-api/internal/commons/types"

type AssetRate struct {
	Name     string
	Symbol   types.CoinSymbol
	Value    float64
	Currency types.CurrencySymbol
}
