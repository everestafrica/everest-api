package models

import "github.com/everestafrica/everest-api/internal/commons/types"

type AccountTransaction struct {
	Base
	UserId        string                    `json:"user_id"`
	MonoId        *string                   `json:"mono_id"`
	TransactionId string                    `json:"transaction_id"`
	Institution   string                    `json:"institution"`
	Currency      string                    `json:"currency"`
	Amount        int                       `json:"amount"`
	Balance       int                       `json:"balance"`
	Date          string                    `json:"date"`
	Narration     string                    `json:"narration"`
	Type          types.TransactionType     `json:"type"`
	Category      types.TransactionCategory `json:"category"`
}
