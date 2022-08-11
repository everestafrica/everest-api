package models

import "github.com/everestafrica/everest-api/internal/commons/types"

type Budget struct {
	Base
	TotalValue string `json:"total_value"`
	Categories []struct {
		Name   types.TransactionCategory `json:"name"`
		Amount int                       `json:"amount"`
	} `json:"categories"`
}
