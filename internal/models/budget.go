package models

import (
	"github.com/everestafrica/everest-api/internal/commons/types"
)

type Budget struct {
	Base
	UserId      string     `json:"user_id"`
	TotalAmount int        `json:"total_amount"`
	Categories  []Category `json:"categories"`
	StartDate   string     `json:"start_date"`
	EndDate     string     `json:"end_date"`
}

type Category struct {
	Base
	BudgetId int                       `json:"budget_id"`
	Name     types.TransactionCategory `json:"name"`
	Amount   int                       `json:"amount" gorm:"default:0"`
}
