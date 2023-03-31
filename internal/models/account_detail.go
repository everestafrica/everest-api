package models

import "github.com/everestafrica/everest-api/internal/commons/types"

type AccountDetail struct {
	Base
	UserId          string                `json:"user_id"`
	AccountId       string                `json:"account_id"`
	Institution     string                `json:"institution"`
	InstitutionType types.InstitutionType `json:"institution_type"`
	AccountNumber   string                `json:"account_number"`
	Balance         int                   `json:"balance"`
	Currency        string                `json:"currency"`
}
