package models

import "time"

type AccountDetail struct {
	Base
	UserId        string `json:"user_id"`
	AccountId     string `json:"account_id"`
	MonoId        string `json:"mono_id"`
	Institution   string `json:"institution"`
	AccountNumber string `json:"account_number"`
	Balance       int    `json:"balance"`
	Currency      string `json:"currency"`
}

type UserMono struct {
	UserId string `json:"user_id"`
	MonoId string `json:"mono_id"`
}

type Webhook struct {
	Event string      `json:"event"`
	Data  WebhookData `json:"data"`
}
type WebhookData struct {
	Meta struct {
		Status string `json:"status"`
	} `json:"meta"`
	Account AccountDetail `json:"account"`
}
type WebhookResult struct {
	Test      string    `json:"test"`
	CreatedAt time.Time `json:"created_at"`
}
