package models

import "time"

type CreateTournamentRequest struct {
	Name      string    `json:"name"`
	PrizePool float32   `json:"prize_pool"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
}
