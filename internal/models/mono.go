package models

// Contains the model struct for everything that involves Mono
import "time"

type Balance struct {
	AccountID     uint   `json:"account_id"`
	AccountNumber string `json:"account_number"`
	Amount        int    `json:"amount"`
	Currency      string `json:"currency"`
}

type AccountDetail struct {
	Base
	MonoID      string  `json:"mono_id"`
	UserID      string  `json:"user_id" gorm:"foreignKey:user_id"`
	Institution string  `json:"institution"`
	Balance     Balance `json:"balance" gorm:"foreignKey:account_id"`
}

type Transaction struct {
	Base
	UserID  string `json:"user_id" gorm:"foreignKey:user_id" json:"user_id,omitempty"`
	Paging  `json:"paging"`
	TxnData []TxnData `json:"txn_data" json:"txn_data,omitempty"`
}
type Paging struct {
	Total    string `json:"total"`
	Page     string `json:"page"`
	Previous string `json:"previous"`
	Next     string `json:"next"`
}

type TxnData struct {
	Id            string `json:"_id"`
	TransactionID string `json:"transaction_id" gorm:"foreignKey:transaction_id"`
	Amount        int    `json:"amount"`
	Balance       int    `json:"balance"`
	Date          string `json:"date"`
	Narration     string `json:"narration"`
	Type          string `json:"type"`
	Category      string `json:"category"`
}

type Webhook struct {
	Event string      `json:"event"`
	Data  WebhookData `json:"data"`
}
type WebhookData struct {
	Meta    WebhookMetadata `json:"meta"`
	Account AccountDetail   `json:"account"`
}
type WebhookMetadata struct {
	Status string `json:"status"`
}
type WebhookResult struct {
	Test      string    `json:"test"`
	CreatedAt time.Time `json:"created_at"`
}
