package tournamentService

import (
	"VykingProject/models"
	repo "VykingProject/repository/prize"
	"context"

	"github.com/google/uuid"
)

type PrizeService interface {
	CreatePrize(ctx context.Context, prize models.CreatePrizeRequest) error
	DistributePrizes(ctx context.Context, tournamentID uuid.UUID) error
}

type prizeService struct {
	repo repo.PrizeRepository
}

func NewPrizeService(repo repo.PrizeRepository) PrizeService {
	return &prizeService{repo: repo}
}

func (s *prizeService) CreatePrize(ctx context.Context, request models.CreatePrizeRequest) error {
	var prize models.Prize
	prize.ID = uuid.New()
	prize.PlayerId = request.PlayerId
	prize.TournamentId = request.TournamentId
	prize.Placement = request.Placement
	err := s.repo.CreatePrize(ctx, prize)
	return err
}

func (s *prizeService) DistributePrizes(ctx context.Context, tournamentID uuid.UUID) error {
	return s.repo.DistributePrizes(ctx, tournamentID)
}
