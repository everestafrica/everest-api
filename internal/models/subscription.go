package models

import "time"

type Subscription struct {
	Base
	Product     string    `json:"product"`
	Price       string    `json:"price"`
	Currency    string    `json:"currency"`
	Logo        string    `json:"logo"`
	Frequency   string    `json:"frequency"`
	NextPayment time.Time `json:"next_payment"`
}
