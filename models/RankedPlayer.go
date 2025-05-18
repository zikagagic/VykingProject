package models

import (
	"github.com/google/uuid"
)

type RankedPlayer struct {
	ID             uuid.UUID `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	AccountBalance float32   `json:"accountBalance"`
	Rank           int32     `json:"rank"`
}
