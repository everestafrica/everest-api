package models

import (
	"github.com/everestafrica/everest-api/internal/commons/types"
)

type Budget struct {
	Base
	BudgetId    string     `json:"budget_id"`
	UserId      string     `json:"user_id"`
	TotalAmount int        `json:"total_amount"`
	Categories  []Category `json:"categories"`
	StartDay    string     `json:"start_day"`
	EndDay      string     `json:"end_day"`
}

type Category struct {
	Base
	BudgetId string                    `json:"budget_id"`
	Name     types.TransactionCategory `json:"name"`
	Amount   int                       `json:"amount"`
}
