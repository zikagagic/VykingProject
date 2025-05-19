package repository

import (
	"VykingProject/models"
	"context"

	"github.com/google/uuid"
)

type TournamentRepository interface {
	CreateTournament(ctx context.Context, tournament models.Tournament) error
	GetAllTournaments(ctx context.Context) ([]models.Tournament, error)
	GetTournamentByID(ctx context.Context, tournamentID uuid.UUID) (models.Tournament, error)
	DeleteTournament(ctx context.Context, tournamentID uuid.UUID) error
}
