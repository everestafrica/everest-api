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
	FirstName   string `json:"first_name" validate:"required"`
	LastName    string `json:"last_name" validate:"required"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email" validate:"required"`
	Password    string `json:"password" validate:"required"`
}
type RegisterResponse struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
type LoginResponse struct {
	Token     *TokenResponse `json:"token"`
	FirstName string         `json:"first_name"`
	LastName  string         `json:"last_name"`
	Email     string         `json:"email"`
}
type SendCodeRequest struct {
	IsEmail  bool   `json:"is_email" form:"is_email"`
	Receiver string `json:"receiver" form:"receiver" binding:"required"`
}

// Claims represent the structure of the JWT token
type Claims struct {
	Email string `json:"email"`
	ID    string `json:"id"`
	jwt.StandardClaims
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresAt   int64  `json:"expires_at"`
	Issuer      string `json:"issuer"`
}

type SubscriptionRequest struct {
	Product     string         `json:"product" validate:"required"`
	Price       string         `json:"price" validate:"required"`
	Currency    CurrencySymbol `json:"currency" validate:"required"`
	Logo        string         `json:"logo"`
	Frequency   Frequency      `json:"frequency" validate:"required"`
	NextPayment time.Time      `json:"next_payment" validate:"required"`
}

type CreateBudgetRequest struct {
	TotalAmount int          `json:"total_amount" validate:"required"`
	Categories  []Categories `json:"categories"  validate:"required"`
	Month       string       `json:"month" validate:"required"`
	Year        int          `json:"year" validate:"required"`
}
type Categories struct {
	Name   TransactionCategory `json:"name"`
	Amount int                 `json:"amount"`
}

type UpdateBudgetRequest struct {
	Categories []Categories `json:"categories"  validate:"required"`
	Month      string       `json:"month" validate:"required"`
	Year       int          `json:"year" validate:"required"`
}
type CryptoWalletRequest struct {
	Address string `json:"address"`
	Symbol  string `json:"symbol"`
}

type Pagination struct {
	Number int
	Size   int
}

type TxnFlowResponse struct {
	Total     float64 `json:"total"`
	DateRange string  `json:"date_range"`
}

type DateRange struct {
	From time.Time
	To   time.Time
}

type (
	CryptoSymbol        string
	CryptoName          string
	CurrencySymbol      string
	Frequency           string
	TransactionType     string
	TransactionCategory string
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
