package models

import (
	"github.com/google/uuid"
)

type Prize struct {
	ID           uuid.UUID `json:"id"`
	PlayerId     uuid.UUID `json:"playerId"`
	TournamentId uuid.UUID `json:"tournamentId"`
	Placement    int       `json:"placement"`
}
