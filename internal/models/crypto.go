package models

type CryptoDetail struct {
	Base
	UserId        string `json:"user_id"`
	WalletAddress string `json:"wallet_address"`
	Balance       int    `json:"balance"`
	CoinName      string `json:"coin_name"`
	Symbol        string `json:"symbol"`
}

type CryptoTransaction struct {
	Base
	UserId   string `json:"user_id"`
	Hash     string `json:"hash"`
	CoinName string `json:"coin_name"`
	Symbol   string `json:"symbol"`
	Amount   int    `json:"amount"`
	Value    int    `json:"value"`
	Date     string `json:"date"`
	Type     string `json:"type"`
}
