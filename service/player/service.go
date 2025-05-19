package playerService

import (
	"VykingProject/models"
	repo "VykingProject/repository/player"
	"context"

	"github.com/google/uuid"
)

type PlayerService interface {
	CreatePlayer(ctx context.Context, request models.CreatePlayerRequest) (models.Player, error)
	GetAllPlayers(ctx context.Context) ([]models.Player, error)
	GetPlayerByID(ctx context.Context, id uuid.UUID) (models.Player, error)
	GetPlayersByRanking(ctx context.Context) ([]models.RankedPlayer, error)
	DeletePlayer(ctx context.Context, id uuid.UUID) error
}

type playerService struct {
	repo repo.PlayerRepository
}

func NewPlayerService(repo repo.PlayerRepository) PlayerService {
	return &playerService{repo: repo}
}

func (s *playerService) CreatePlayer(ctx context.Context, request models.CreatePlayerRequest) (models.Player, error) {
	var player models.Player
	player.ID = uuid.New()
	player.Name = request.Name
	player.Email = request.Email
	player.AccountBalance = request.AccountBalance
	err := s.repo.CreatePlayer(ctx, player)
	return player, err
}

func (s *playerService) GetAllPlayers(ctx context.Context) ([]models.Player, error) {
	return s.repo.GetAllPlayers(ctx)
}

func (s *playerService) GetPlayersByRanking(ctx context.Context) ([]models.RankedPlayer, error) {
	players, err := s.repo.GetPlayersByRanking(ctx)
	return players, err
}

func (s *playerService) GetPlayerByID(ctx context.Context, id uuid.UUID) (models.Player, error) {
	return s.repo.GetPlayerByID(ctx, id)
}

func (s *playerService) DeletePlayer(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeletePlayer(ctx, id)
}
