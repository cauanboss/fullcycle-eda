package dto

import "time"

type CreateBalanceInput struct {
	AccountID string  `json:"account_id"`
	Balance   float64 `json:"balance"`
}

type CreateBalanceOutput struct {
	ID        string    `json:"id"`
	AccountID string    `json:"account_id"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
