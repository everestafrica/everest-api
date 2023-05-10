package model

// This table is used to track user activities

type Tracker struct {
	UserId              string `json:"user_id"`
	VerifiedEmail       bool   `json:"verified_email"`
	CreatedBudget       bool   `json:"created_budget"`
	CreatedSubscription bool   `json:"created_subscription"`
	CreatedDebt         bool   `json:"created_debt"`
	AllowRanking        bool   `json:"allow_ranking"`
}

type Level string

const (
	LevelOne   Level = "free"
	LevelTwo   Level = "pro"
	LevelThree Level = "premium"
)
