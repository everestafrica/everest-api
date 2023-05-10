package model

import "github.com/everestafrica/everest-api/internal/commons/types"

type Category struct {
	Base
	UserId  string `json:"-"`
	Name    string `json:"name"`
	Icon    string `json:"icon"`
	Unicode string `json:"unicode"`
	Custom  bool   `json:"custom"`
}

type PriceAlert struct {
	Base
	UserId   string  `json:"user_id"`
	Asset    string  `json:"asset"`
	IsCrypto bool    `json:"is_crypto"`
	Target   float64 `json:"target"`
}

type NewsInterest struct {
	Base
	UserId   string
	Interest types.NewsInterest
}
