package types

import (
	"github.com/golang-jwt/jwt"
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
	Code string `json:"code"`
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
		ID        string `json:"_id"`
		Type      string `json:"type"`
		Amount    int    `json:"amount"`
		Narration string `json:"narration"`
		Date      string `json:"date"`
		Balance   int    `json:"balance"`
		Currency  string `json:"currency"`
		Category  string `json:"category"`
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
