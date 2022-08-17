package types

import (
	"github.com/golang-jwt/jwt"
	"time"
)

type GenericResponse struct {
	Success bool        `json:"success"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type RegisterRequest struct {
	FirstName   string `json:"first_name" binding:"required"`
	LastName    string `json:"last_name" binding:"required"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email" binding:"required"`
	Password    string `json:"password" binding:"required"`
}

type RegisterResponse struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type LoginResponse struct {
	Token     *TokenResponse `json:"token"`
	FirstName string         `json:"first_name"`
	LastName  string         `json:"last_name"`
	Email     string         `json:"email"`
}

// Claims represent the structure of the JWT token
type Claims struct {
	Email string `json:"email"`
	ID    string `json:"id"`
	jwt.StandardClaims
}

type MonoAccountIdRequest struct {
	Code string `json:"code" binding:"required"`
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresAt   int64  `json:"expires_at"`
	Issuer      string `json:"issuer"`
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
type SubscriptionRequest struct {
	Product     string         `json:"product" binding:"required"`
	Price       string         `json:"price" binding:"required"`
	Currency    CurrencySymbol `json:"currency" binding:"required"`
	Logo        string         `json:"logo" binding:"required"`
	Frequency   Frequency      `json:"frequency" binding:"required"`
	NextPayment time.Time      `json:"next_payment" binding:"required"`
}

type BudgetRequest struct {
	Categories []struct {
		Name   TransactionCategory `json:"name"`
		Amount int                 `json:"amount"`
	} `json:"categories"`
}
type CryptoSymbol string
type CryptoName string
type CurrencySymbol string
type Frequency string
type TransactionType string
type TransactionCategory string

const (
	NGN CurrencySymbol = "NGN"
	GBP CurrencySymbol = "GBP"
	USD CurrencySymbol = "USD"
	EUR CurrencySymbol = "EUR"
	KES CurrencySymbol = "KES"
	GHC CurrencySymbol = "GHC"
	ZAR CurrencySymbol = "ZAR"

	BTC CryptoSymbol = "BTC"
	ETH CryptoSymbol = "ETH"
	BNB CryptoSymbol = "BNB"

	MONTHLY  Frequency = "MONTHLY"
	ANNUALLY Frequency = "ANNUALLY"
)

const (
	Credit TransactionType = "credit"
	Debit  TransactionType = "debit"

	Grocery             TransactionCategory = "groceries"
	Entertainment       TransactionCategory = "entertainment"
	Travel              TransactionCategory = "travel"
	Transport           TransactionCategory = "transportation"
	Salary              TransactionCategory = "salary"
	Investment          TransactionCategory = "investment"
	PhoneAndInternet    TransactionCategory = "phone_and_internet"
	Food                TransactionCategory = "food"
	Health              TransactionCategory = "health"
	SelfCare            TransactionCategory = "self_care"
	LoanRepayment       TransactionCategory = "loan_repayment"
	Bills               TransactionCategory = "bills/fees"
	Transfer            TransactionCategory = "transfer"
	OnlineTransactions  TransactionCategory = "online_transactions"
	OfflineTransactions TransactionCategory = "offline_transactions"
	BankCharges         TransactionCategory = "bank_charges"
	AtmWithdrawal       TransactionCategory = "atm_withdrawal"
	Miscellaneous       TransactionCategory = "miscellaneous"
	Others              TransactionCategory = "others"
)

// Track debtors, snap receipt, family tracker
// GiftsAndDonations TransactionCategory = "gifts_and_donations" 	Education TransactionCategory = "education"
