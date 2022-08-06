package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	Base
	FirstName   string  `json:"first_name"`
	LastName    string  `json:"last_name"`
	Email       string  `json:"email" gorm:"unique"`
	PhoneNumber string  `json:"phone_number"`
	UserId      string  `json:"user_id" gorm:"unique;not null;type:varchar(100)"`
	DOB         *string `json:"dob"`
	Country     *string `json:"country"`
	Currency    *string `json:"currency"`
	Password    string  `json:"-"`
	MonoId      *string `json:"mono_id"`
	MonoCode    *string `json:"mono_code"`
	MonoStatus  bool    `json:"mono_status"`
	MonoReauth  *string `json:"mono_reauth"`
	NetWorth    *int    `json:"net_worth"`
	Persona     *string `json:"persona"`
}

// Base contains common columns for all tables
type Base struct {
	ID uint `gorm:"primaryKey" json:"id"`
	//UUID      string    `json:"uuid"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// BeforeCreate will set User struct before every insert
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.UserId = uuid.NewString()
	return nil
}

//func (b *Base) BeforeSave(tx *gorm.DB) (err error) {
//	b.UUID = uuid.NewString()
//	return nil
//}
