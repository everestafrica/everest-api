package models

type AccountDetail struct {
	Base
	UserId        string `json:"user_id"`
	AccountId     string `json:"account_id"`
	Institution   string `json:"institution"`
	AccountNumber string `json:"account_number"`
	Balance       int    `json:"balance"`
	Currency      string `json:"currency"`
}
