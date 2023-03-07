package models

import (
	"github.com/everestafrica/everest-api/internal/commons/types"
	"time"
)

type AccountTransaction struct {
	Base
	UserId        string                    `json:"user_id"`
	AccountId     *string                   `json:"account_id"`
	TransactionId string                    `json:"transaction_id"`
	Institution   string                    `json:"institution"`
	Currency      types.CurrencySymbol      `json:"currency"`
	Amount        float64                   `json:"amount"`
	Balance       *float64                  `json:"balance"`
	Narration     string                    `json:"narration"`
	Merchant      string                    `json:"merchant"`
	IsRecurring   bool                      `json:"is_recurring"`
	Type          types.TransactionType     `json:"type"`
	Category      types.TransactionCategory `json:"category"`
	Date          time.Time                 `json:"date"`
}
