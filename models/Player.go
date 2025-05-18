package models

import (
	"github.com/google/uuid"
)

type Player struct {
	ID             uuid.UUID `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	AccountBalance float32   `json:"accountBalance"`
}
