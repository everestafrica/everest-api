package models

import (
	"github.com/everestafrica/everest-api/internal/commons/types"
	"time"
)

type Subscription struct {
	Base
	UserId      string                 `json:"user_id"`
	Product     string                 `json:"product"`
	Price       float64                `json:"price"`
	Currency    types.CurrencySymbol   `json:"currency"`
	Icon        string                 `json:"icon"`
	Frequency   types.Frequency        `json:"frequency"`
	Type        types.SubscriptionType `json:"type"`
	NextPayment time.Time              `json:"next_payment"`
}
