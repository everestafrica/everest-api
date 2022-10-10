package models

import (
	"time"
)

type MonoUser struct {
	Base
	UserId      string    `json:"user_id" gorm:"unique;not null;type:varchar(100)"`
	MonoId      string    `json:"mono_id"`
	MonoStatus  string    `json:"mono_status"`
	LastRefresh time.Time `json:"last_refresh"`
	Reauth      bool      `json:"reauth" gorm:"default:true"`
	ReauthToken *string   `json:"reauth_token"`
}
