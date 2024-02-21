package models

import (
	"github.com/everestafrica/everest-api/internal/commons/types"
	"github.com/everestafrica/everest-api/internal/commons/utils"
	"gorm.io/gorm"
	"time"
)

type CashAccount struct {
	Base
	UserId          string                `json:"user_id"`
	AccountId       string                `json:"account_id"`
	Institution     string                `json:"institution"`
	InstitutionType types.InstitutionType `json:"institution_type"`
	AccountNumber   string                `json:"account_number"`
	Balance         int                   `json:"balance"`
	Currency        string                `json:"currency"`
}

type CashTransaction struct {
	Base
	Id              string                    `json:"id"`
	UserId          string                    `json:"user_id"`
	AccountId       *string                   `json:"account_id"`
	Institution     string                    `json:"institution"`
	InstitutionType types.InstitutionType     `json:"institution_type"`
	Currency        types.CurrencySymbol      `json:"currency"`
	Amount          float64                   `json:"amount"`
	Balance         *float64                  `json:"balance"`
	Narration       string                    `json:"narration"`
	Merchant        string                    `json:"merchant"`
	IsRecurring     bool                      `json:"is_recurring"`
	Type            types.TransactionType     `json:"type"`
	Category        types.TransactionCategory `json:"category"`
	Date            time.Time                 `json:"date"`
}

func (c CashTransaction) BeforeCreate(tx *gorm.DB) (err error) {
	c.Id = utils.GetUUID()
	return err
}
