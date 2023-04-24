package types

import (
	"github.com/golang-jwt/jwt"
	"time"
)

type (
	GenericResponse struct {
		Success bool        `json:"success"`
		Message interface{} `json:"message"`
		Data    interface{} `json:"data,omitempty"`
	}
	RegisterRequest struct {
		FirstName   string `json:"first_name" validate:"required"`
		LastName    string `json:"last_name" validate:"required"`
		PhoneNumber string `json:"phone_number"`
		Email       string `json:"email" validate:"required"`
		Password    string `json:"password" validate:"required"`
		Code        string `json:"code" validate:"required"`
	}
	RegisterResponse struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}
	LoginRequest struct {
		Email    string `json:"email" validate:"required"`
		Password string `json:"password" validate:"required"`
	}
	LoginResponse struct {
		Token     *TokenResponse `json:"token"`
		FirstName string         `json:"first_name"`
		LastName  string         `json:"last_name"`
		Email     string         `json:"-"`
	}
	SendCodeRequest struct {
		IsEmail  bool   `json:"is_email"`
		Receiver string `json:"receiver" validate:"required"`
	}
	// Claims represent the structure of the JWT token
	Claims struct {
		Email string `json:"email"`
		ID    string `json:"id"`
		jwt.StandardClaims
	}
	TokenResponse struct {
		AccessToken string `json:"access_token"`
		ExpiresAt   int64  `json:"expires_at"`
		Issuer      string `json:"issuer"`
	}
	SubscriptionRequest struct {
		Product     string         `json:"product" validate:"required"`
		Price       float64        `json:"price" validate:"required"`
		Currency    CurrencySymbol `json:"currency" validate:"required"`
		Icon        string         `json:"icon"`
		Frequency   Frequency      `json:"frequency" validate:"required"`
		NextPayment string         `json:"next_payment" validate:"required"`
	}
	CreateBudgetRequest struct {
		TotalAmount int          `json:"total_amount" validate:"required"`
		Categories  []Categories `json:"categories"  validate:"required"`
		Month       string       `json:"month" validate:"required"`
		Year        int          `json:"year" validate:"required"`
	}
	Categories struct {
		Name   TransactionCategory `json:"name"`
		Amount int                 `json:"amount"`
	}
	UpdateBudgetRequest struct {
		Categories []Categories `json:"categories"  validate:"required"`
		Month      string       `json:"month" validate:"required"`
		Year       int          `json:"year" validate:"required"`
	}
	CryptoWalletRequest struct {
		Address string `json:"address"`
		Symbol  string `json:"symbol"`
	}
	CreateDebtRequest struct {
		Amount           int64     `json:"amount" validate:"required"`
		Reason           string    `json:"reason" validate:"required"`
		CounterpartyName string    `json:"counterparty_name" validate:"required"`
		Due              time.Time `json:"due" validate:"required"`
	}
	UpdateDebtRequest struct {
		Amount           *int64     `json:"amount"`
		Due              *time.Time `json:"due"`
		CounterpartyName *string    `json:"counterparty_name"`
		Reason           *string    `json:"reason"`
	}
	AddStockRequest struct {
		Symbol string
	}
	UpdateTransactionRequest struct {
		IsRecurring bool                `json:"is_recurring"`
		Category    TransactionCategory `json:"category"`
	}
	CreateTransactionRequest struct {
		Amount        float64             `json:"amount"`
		Narration     string              `json:"narration"`
		Merchant      string              `json:"merchant"`
		Type          TransactionType     `json:"type"`
		Currency      CurrencySymbol      `json:"currency"`
		TransactionId string              `json:"transaction_id"`
		IsRecurring   bool                `json:"is_recurring"`
		Category      TransactionCategory `json:"category"`
		Date          time.Time           `json:"date"`
	}
	CreateCustomCategory struct {
		Name  string `json:"name"`
		Emoji string `json:"emoji"`
	}
	CreatePriceAlert struct {
		Asset    string  `json:"asset"`
		Target   float64 `json:"target"`
		IsCrypto bool    `json:"is_crypto"`
	}
	AddNewsInterest struct {
		Category string `json:"category"`
	}
	Pagination struct {
		Page int
		Size int
	}
	TxnFlowResponse struct {
		Total     float64 `json:"total"`
		DateRange string  `json:"date_range"`
	}
	DateRange struct {
		From time.Time
		To   time.Time
	}
)

type (
	CryptoSymbol        string
	CryptoName          string
	CurrencySymbol      string
	Frequency           string
	SubscriptionType    string
	TransactionType     string
	TransactionCategory string
	DebtType            string
	DebtStatus          string
	NewsInterest        string
)

const (
	NGN CurrencySymbol = "NGN"
	GBP CurrencySymbol = "GBP"
	USD CurrencySymbol = "USD"
	EUR CurrencySymbol = "EUR"
	KES CurrencySymbol = "KES"
	GHC CurrencySymbol = "GHC"
	ZAR CurrencySymbol = "ZAR"

	BTC  CryptoSymbol = "BTC"
	ETH  CryptoSymbol = "ETH"
	BSC  CryptoSymbol = "BSC"
	USDT CryptoSymbol = "USDT"
	SOL  CryptoSymbol = "SOL"
	DOGE CryptoSymbol = "DOGE"

	MONTHLY  Frequency = "MONTHLY"
	ANNUALLY Frequency = "ANNUALLY"
)

const (
	Credit TransactionType = "credit"
	Debit  TransactionType = "debit"

	Media    SubscriptionType = "media"
	Software SubscriptionType = "software"
	Services SubscriptionType = "services"

	DebtStatusPending DebtStatus = "pending"
	DebtStatusPaid    DebtStatus = "paid"
	DebtStatusFailed  DebtStatus = "failed"

	DebtTypeDebtor   DebtType = "debtor"
	DebtTypeCreditor DebtType = "creditor"

	Banking        NewsInterest = "banking"
	Investing      NewsInterest = "investing"
	Saving         NewsInterest = "saving"
	CryptoCurrency NewsInterest = "cryptocurrency"
	Stocks         NewsInterest = "stocks"
	Insurance      NewsInterest = "insurance"
	Business       NewsInterest = "business"
	Fintech        NewsInterest = "fintech"

	Groceries          TransactionCategory = "Groceries"
	Shopping           TransactionCategory = "Shopping"
	Utilities          TransactionCategory = "Utilities and Bills"
	Housing            TransactionCategory = "Housing"
	Entertainment      TransactionCategory = "Entertainment"
	Travel             TransactionCategory = "Travel"
	Delivery           TransactionCategory = "Delivery"
	Transportation     TransactionCategory = "Transportation"
	Income             TransactionCategory = "Income"
	Investment         TransactionCategory = "Investment"
	PhoneAndInternet   TransactionCategory = "Phone & Internet"
	Food               TransactionCategory = "Food"
	Healthcare         TransactionCategory = "Healthcare"
	LoanRepayment      TransactionCategory = "Loan Repayment"
	LoanOut            TransactionCategory = "Loan Out"
	Transfer           TransactionCategory = "Transfer"
	OnlineTransaction  TransactionCategory = "Online Transaction"
	OfflineTransaction TransactionCategory = "Offline Transaction"
	BankCharges        TransactionCategory = "Bank Charges"
	AtmWithdrawal      TransactionCategory = "ATM Withdrawal"
	Miscellaneous      TransactionCategory = "Miscellaneous"
	GiftsAndDonations  TransactionCategory = "Gifts & Donations"
	Education          TransactionCategory = "Education"

	Others TransactionCategory = "others"
)

// Track debtors, snap receipt, family tracker

//["Groceries",  "Shopping",  "Utilities and Bills",  "Housing",  "Entertainment",  "Travel",  "Delivery",  "Transportation",  "Income",  "Investment",  "Phone & Internet",  "Food",  "Healthcare",  "Loan Repayment",  "Loan Out",  "Transfer",  "Online Transaction",  "Offline Transaction",  "Bank Charges",  "ATM Withdrawal",  "Miscellaneous",  "Gifts & Donations",  "Education"]
// trending  - places, items
