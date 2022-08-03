package models

// Contains the model struct for everything that involves Mono
import "time"

type Balance struct {
	AccountId     uint   `json:"account_id"`
	AccountNumber string `json:"account_number"`
	Amount        int    `json:"amount"`
	Currency      string `json:"currency"`
}

type AccountDetail struct {
	Base
	MonoId      string  `json:"mono_id"`
	UserId      string  `json:"user_id" gorm:"foreignKey:user_id"`
	Institution string  `json:"institution"`
	Balance     Balance `json:"balance" gorm:"foreignKey:account_id"`
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
