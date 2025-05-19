package repository

import (
	"VykingProject/models"
	"context"

	"github.com/google/uuid"
)

type PrizeRepository interface {
	CreatePrize(ctx context.Context, prize models.Prize) error
	DistributePrizes(ctx context.Context, tournamentID uuid.UUID) error
}
