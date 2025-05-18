package models

type CreatePlayerRequest struct {
	Name           string  `json:"name"`
	Email          string  `json:"email"`
	AccountBalance float32 `json:"account_balance"`
}
