package models

// This table is used to track user activities

type Tracker struct {
	UserId              string `json:"user_id"`
	VerifiedEmail       bool   `json:"verified_email"`
	CreatedBudget       bool   `json:"created_budget"`
	CreatedSubscription bool   `json:"created_subscription"`
	AllowRanking        bool   `json:"allow_ranking"`
}
