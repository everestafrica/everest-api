package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	Base
	UserId      string  `json:"user_id" gorm:"unique;not null;type:varchar(100)"`
	FirstName   string  `json:"first_name"`
	LastName    string  `json:"last_name"`
	Email       string  `json:"email" gorm:"unique"`
	Username    string  `json:"username" gorm:"unique"`
	PhoneNumber string  `json:"phone_number"`
	DOB         *string `json:"dob"`
	Country     *string `json:"country"`
	Currency    *string `json:"currency"`
	Password    string  `json:"-"`
	NetWorth    *int    `json:"net_worth"`
	Persona     *string `json:"persona"`
	Avatar      *string `json:"avatar"`
	PushToken   *string `json:"push_token"`
}

// Base contains common columns for all tables
type Base struct {
	ID int `gorm:"primaryKey" json:"id"`
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
//	return nil uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
//}
