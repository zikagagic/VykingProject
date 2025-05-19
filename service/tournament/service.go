package tournamentService

import (
	"VykingProject/models"
	repo "VykingProject/repository/tournament"
	"context"

	"github.com/google/uuid"
)

type TournamentService interface {
	CreateTournament(ctx context.Context, tournament models.CreateTournamentRequest) (models.Tournament, error)
	GetAllTournaments(ctx context.Context) ([]models.Tournament, error)
	GetTournamentByID(ctx context.Context, id uuid.UUID) (models.Tournament, error)
	DeleteTournament(ctx context.Context, id uuid.UUID) error
}

type tournamentService struct {
	repo repo.TournamentRepository
}

func NewTournamentService(repo repo.TournamentRepository) TournamentService {
	return &tournamentService{repo: repo}
}

func (s *tournamentService) CreateTournament(ctx context.Context, t models.CreateTournamentRequest) (models.Tournament, error) {
	var tournament models.Tournament

	tournament.ID = uuid.New()
	tournament.Name = t.Name
	tournament.PrizePool = t.PrizePool
	tournament.StartDate = t.StartDate
	tournament.EndDate = t.EndDate
	err := s.repo.CreateTournament(ctx, tournament)
	return tournament, err
}

func (s *tournamentService) GetAllTournaments(ctx context.Context) ([]models.Tournament, error) {
	return s.repo.GetAllTournaments(ctx)
}

func (s *tournamentService) GetTournamentByID(ctx context.Context, id uuid.UUID) (models.Tournament, error) {
	return s.repo.GetTournamentByID(ctx, id)
}

func (s *tournamentService) DeleteTournament(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteTournament(ctx, id)
}
