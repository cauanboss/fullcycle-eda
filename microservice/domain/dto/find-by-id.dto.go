package dto

type FindByIdInput struct {
	ID string `json:"id"`
}

type FindByIdOutput struct {
	ID        string  `json:"id"`
	AccountID string  `json:"account_id"`
	Balance   float64 `json:"balance"`
}
