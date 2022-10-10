package models

import (
	"github.com/everestafrica/everest-api/internal/commons/types"
	"time"
)

type AccountTransaction struct {
	Base
	UserId        string                    `json:"user_id"`
	MonoId        *string                   `json:"mono_id"`
	TransactionId string                    `json:"transaction_id"`
	Institution   string                    `json:"institution"`
	Currency      string                    `json:"currency"`
	Amount        float64                   `json:"amount"`
	Balance       float64                   `json:"balance"`
	Narration     string                    `json:"narration"`
	Type          types.TransactionType     `json:"type"`
	Category      types.TransactionCategory `json:"category"`
	Date          time.Time                 `json:"date"`
}
