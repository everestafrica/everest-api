package models

import (
	"github.com/everestafrica/everest-api/internal/commons/types"
)

type Budget struct {
	Base
	UserId   string                    `json:"-"`
	BudgetId string                    `json:"budget_id"`
	Category types.TransactionCategory `json:"category"`
	Amount   int                       `json:"amount" gorm:"default:0"`
	Month    string                    `json:"-"`
	Year     int                       `json:"-"`
}
