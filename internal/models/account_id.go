package models

type AccountId struct {
	Base
	UserId string `json:"user_id"`
	MonoId string `json:"mono_id"`
}
