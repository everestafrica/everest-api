package models

type Asset struct {
	Base
	UserId   string  `json:"user_id"`
	Symbol   string  `json:"symbol"`
	Name     string  `json:"name"`
	Image    string  `json:"image"`
	IsCrypto string  `json:"is_crypto"`
	Value    float64 `json:"value"`
}
