package repository

import (
	"VykingProject/models"
	"context"

	"github.com/google/uuid"
)

type PlayerRepository interface {
	CreatePlayer(ctx context.Context, player models.Player) error
	GetAllPlayers(ctx context.Context) ([]models.Player, error)
	GetPlayersByRanking(ctx context.Context) ([]models.RankedPlayer, error)
	GetPlayerByID(ctx context.Context, playerID uuid.UUID) (models.Player, error)
	DeletePlayer(ctx context.Context, playerID uuid.UUID) error
}
