package models

import "github.com/google/uuid"

type CreatePrizeRequest struct {
	PlayerId     uuid.UUID `json:"playerId"`
	TournamentId uuid.UUID `json:"tournamentId"`
	Placement    int       `json:"placement"`
}
