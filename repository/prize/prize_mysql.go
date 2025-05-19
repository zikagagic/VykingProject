package repository

import (
	"VykingProject/models"
	"context"
	"database/sql"

	"github.com/google/uuid"
)

type MYSQLPrizeRepository struct {
	db *sql.DB
}

func NewMYSQLTournamentRepository(db *sql.DB) *MYSQLPrizeRepository {
	return &MYSQLPrizeRepository{db: db}
}

func (r *MYSQLPrizeRepository) CreatePrize(ctx context.Context, prize models.Prize) error {
	query := `INSERT INTO prize_placement (id, playerID, tournamentID, placement) VALUES (?, ?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, prize.ID, prize.PlayerId, prize.TournamentId, prize.Placement)
	return err
}

func (r *MYSQLPrizeRepository) DistributePrizes(ctx context.Context, tournamentId uuid.UUID) error {
	querry := `CALL distribute_prizes(?)`
	_, err := r.db.ExecContext(ctx, querry, tournamentId)
	return err
}
