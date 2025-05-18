package models

import (
	"time"

	"github.com/google/uuid"
)

type Tournament struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	PrizePool float32   `json:"prizePool"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
}
