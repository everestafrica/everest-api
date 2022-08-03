package models

type Transaction struct {
	Base
	UserId  string `json:"user_id"`
	Paging  `json:"paging"`
	TxnData []TxnData `json:"transaction_data"`
}

type Paging struct {
	Total    string `json:"total"`
	Page     string `json:"page"`
	Previous string `json:"previous"`
	Next     string `json:"next"`
}

type TxnData struct {
	Id            string `json:"_id"`
	TransactionID string `json:"transaction_id"`
	Amount        int    `json:"amount"`
	Balance       int    `json:"balance"`
	Date          string `json:"date"`
	Narration     string `json:"narration"`
	Type          string `json:"type"`
	Category      string `json:"category"`
}
