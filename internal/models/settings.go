package models

type CustomCategory struct {
	Base
	UserId string `json:"-"`
	Name   string `json:"name"`
	Emoji  string `json:"emoji"`
}

type PriceAlert struct {
	Base
	UserId   string  `json:"user_id"`
	Asset    string  `json:"asset"`
	IsCrypto bool    `json:"is_crypto"`
	Target   float64 `json:"target"`
}
