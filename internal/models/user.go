package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	Base
	Name       string  `json:"name" gorm:"unique"`
	Email      string  `json:"email" gorm:"unique"`
	DOB        *string `json:"dob"`
	Country    *string `json:"country"`
	Currency   *string `json:"currency"`
	Password   string  `json:"-"`
	MonoId     *string `json:"mono_id"`
	MonoCode   *string `json:"mono_code"`
	MonoStatus bool    `json:"mono_status"`
	MonoReauth *string `json:"mono_reauth"`
	//AccountDetail *AccountDetail  `json:"account_info"`
	//Crypto        *Crypto         `json:"crypto"`
	//Transactions  []*Transaction  `json:"transactions"`
	//Subscriptions []*Subscription `json:"subscriptions"`
	//Budget        []*Budget       `json:"budgets"`
	//Settings      Settings        `json:"settings"`
	Networth *int    `json:"networth"`
	Persona  *string `json:"persona"`
}

// GenerateISOString generates a time string equivalent to Date.now().toISOString in JavaScript
func GenerateISOString() string {
	return time.Now().UTC().Format("2006-01-02T15:04:05.999Z07:00")
}

// Base contains common columns for all tables
type Base struct {
	ID        uint      `gorm:"primaryKey"`
	UUID      uuid.UUID `json:"_id" gorm:"primaryKey;autoIncrement:false"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
}

// BeforeCreate will set Base struct before every insert
func (base *Base) BeforeCreate(tx *gorm.DB) error {
	// uuid.New() creates a new random UUID or panics.
	base.UUID = uuid.New()

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
