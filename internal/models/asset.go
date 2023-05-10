package models

type Asset struct {
	Base
	UserId   string  `json:"user_id"`
	Symbol   string  `json:"symbol"`
	Name     string  `json:"name"`
	Image    string  `json:"image"`
	Price    float64 `json:"price"`
	Amount   float64 `json:"amount"`
	Value    float64 `json:"value"`
	IsCrypto bool    `json:"is_crypto"`
}

type Stock struct {
	Base
	Name   string `json:"name"   gorm:"unique,not null"`
	Image  string `json:"image"  gorm:"unique,not null"`
	Symbol string `json:"symbol" gorm:"unique,not null"`
}
