package models

type Asset struct {
	Base
	UserId   string  `json:"user_id"`
	Symbol   string  `json:"symbol"`
	Name     string  `json:"name"`
	Image    string  `json:"image"`
	IsCrypto string  `json:"is_crypto"`
	Price    float64 `json:"price"`
	Amount   float64 `json:"amount"`
	Value    float64 `json:"value"`
}

type Stock struct {
	Base
	Name   string `json:"name"   gorm:"unique,not null"`
	Image  string `json:"image"  gorm:"unique,not null"`
	Symbol string `json:"symbol" gorm:"unique,not null"`
}
