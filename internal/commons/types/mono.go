package types

import "time"

type MonoAccountIdRequest struct {
	Code string `json:"code" validate:"required"`
}

type MonoAccountIdResponse struct {
	Id string `json:"id"`
}

type MonoTransactionResponse struct {
	Paging struct {
		Total    int    `json:"total"`
		Page     int    `json:"page"`
		Previous string `json:"previous"`
		Next     string `json:"next"`
	} `json:"paging"`
	Data []struct {
		ID        string          `json:"_id"`
		Type      string          `json:"type"`
		Amount    int             `json:"amount"`
		Narration string          `json:"narration"`
		Date      time.Time       `json:"date"`
		Balance   int             `json:"balance"`
		Currency  string          `json:"currency"`
		Category  TransactionType `json:"category"`
	} `json:"data"`
}

type MonoAccountResponse struct {
	Meta struct {
		DataStatus string `json:"data_status"`
		AuthMethod string `json:"auth_method"`
	} `json:"meta"`
	Account struct {
		ID          string `json:"_id"`
		Institution struct {
			Name     string `json:"name"`
			BankCode string `json:"bankCode"`
			Type     string `json:"type"`
		} `json:"institution"`
		Name          string `json:"name"`
		AccountNumber string `json:"accountNumber"`
		Type          string `json:"type"`
		Balance       int    `json:"balance"`
		Currency      string `json:"currency"`
		Bvn           string `json:"bvn"`
	} `json:"account"`
}
type MonoManualsyncResponse struct {
	Status     string `json:"status"`
	HasNewData *bool  `json:"hasNewData"`
	Code       string `json:"code"`
}
type MonoReauthResponse struct {
	Token string `json:"token"`
}

type MonoWebhookPayload struct {
	Event string      `json:"event"`
	Data  interface{} `json:"data"`
}
type MonoWebhookDataSync struct {
	Account struct {
		Id string `json:"_id"`
	}
}
type MonoWebhookConnect struct {
	Id string `json:"id"`
}
type MonoWebhookUnlink struct {
	Account struct {
		Id string `json:"id"`
	}
}

type MonoWebhookAccountSync struct {
	Account string `json:"account"`
	Updated bool   `json:"updated"`
}

type AccountType string
type InstitutionType string

const (
	SavingsAccount    InstitutionType = "SAVINGS_ACCOUNT"
	InvestmentAccount InstitutionType = "INVESTMENT_ACCOUNT"
	WalletAccount     InstitutionType = "WALLET_ACCOUNT"
	CashAccount       InstitutionType = "CASH_ACCOUNT"
)

const (
	PersonalBanking AccountType = "PERSONAL_BANKING"
	BusinessBanking AccountType = "BUSINESS_BANKING"
)
