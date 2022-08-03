package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Base
	FirstName  string  `json:"first_name"`
	LastName   string  `json:"last_name"`
	Email      string  `json:"email" gorm:"unique"`
	UserId     string  `json:"user_id"`
	DOB        *string `json:"dob"`
	Country    *string `json:"country"`
	Currency   *string `json:"currency"`
	Password   string  `json:"-"`
	MonoId     *string `json:"mono_id"`
	MonoCode   *string `json:"mono_code"`
	MonoStatus bool    `json:"mono_status"`
	MonoReauth *string `json:"mono_reauth"`
	NetWorth   *int    `json:"net_worth"`
	Persona    *string `json:"persona"`
}

// GenerateISOString generates a time string equivalent to Date.now().toISOString in JavaScript
func GenerateISOString() string {
	return time.Now().UTC().Format("2006-01-02T15:04:05.999Z07:00")
}

// Base contains common columns for all tables
type Base struct {
	ID        uint   `gorm:"primaryKey"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// BeforeCreate will set Base struct before every insert
func (base *Base) BeforeCreate(tx *gorm.DB) error {
	// generate timestamps
	t := GenerateISOString()
	base.CreatedAt, base.UpdatedAt = t, t

	return nil
}

// AfterUpdate will update the Base struct after every update
func (base *Base) AfterUpdate(tx *gorm.DB) error {
	// update timestamps
	base.UpdatedAt = GenerateISOString()
	return nil
}
