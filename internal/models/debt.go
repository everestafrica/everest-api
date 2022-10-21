package models

import (
	"github.com/everestafrica/everest-api/internal/commons/types"
	"time"
)

type Debt struct {
	Base
	UserId           string           `json:"user_id" gorm:"index;not null"`
	Amount           int64            `json:"amount" gorm:"not null"`
	CounterpartyName string           `json:"counterparty_name" gorm:"not null"`
	DebtType         types.DebtType   `json:"type" gorm:"not null"`
	Status           types.DebtStatus `json:"status" gorm:"not null,default:pending"`
	Reason           string           `json:"reason"`
	Due              time.Time        `json:"due" gorm:"not null"`
}
