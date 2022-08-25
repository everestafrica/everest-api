package types

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
		Date      string          `json:"date"`
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
			BankCode string `json:"bank_code"`
			Type     string `json:"type"`
		} `json:"institution"`
		Name          string `json:"name"`
		AccountNumber string `json:"account_number"`
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
type MonoWebhookDataSyncData struct {
	Account struct {
		Id string `json:"_id"`
	}
}
type MonoWebhookUnlinkData struct {
	Account struct {
		Id string `json:"id"`
	}
}
