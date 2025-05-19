package repository

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"VykingProject/models"

	"github.com/google/uuid"
)

type MYSQLTournamentRepository struct {
	db *sql.DB
}

func NewMYSQLTournamentRepository(db *sql.DB) *MYSQLTournamentRepository {
	return &MYSQLTournamentRepository{db: db}
}

func (r *MYSQLTournamentRepository) CreateTournament(ctx context.Context, tournament models.Tournament) error {
	query := `INSERT INTO tournament (id, name, prize_pool, start_date, end_date) VALUES (?, ?, ?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, tournament.ID, tournament.Name, tournament.PrizePool, tournament.StartDate, tournament.EndDate)
	return err
}

// GetAllTournaments retrieves all tournaments from the database.
func (r *MYSQLTournamentRepository) GetAllTournaments(ctx context.Context) ([]models.Tournament, error) {
	query := `SELECT id, name, prize_pool, start_date, end_date FROM tournament`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		log.Println("Query error:", err)
		return nil, err
	}
	defer rows.Close()

	var tournaments []models.Tournament
	for rows.Next() {
		var t models.Tournament
		err := rows.Scan(&t.ID, &t.Name, &t.PrizePool, &t.StartDate, &t.EndDate)
		if err != nil {
			return nil, err
		}
		tournaments = append(tournaments, t)
	}
	return tournaments, rows.Err()
}

// GetTournamentByID fetches a specific tournament by ID.
func (r *MYSQLTournamentRepository) GetTournamentByID(ctx context.Context, tournamentID uuid.UUID) (models.Tournament, error) {
	query := `SELECT id, name, prize_pool, start_date, end_date FROM tournament WHERE id = ?`
	var t models.Tournament
	err := r.db.QueryRowContext(ctx, query, tournamentID).Scan(&t.ID, &t.Name, &t.PrizePool, &t.StartDate, &t.EndDate)
	if errors.Is(err, sql.ErrNoRows) {
		return t, errors.New("tournament not found")
	}
	return t, err
}

// DeleteTournament deletes a tournament by ID.
func (r *MYSQLTournamentRepository) DeleteTournament(ctx context.Context, tournamentID uuid.UUID) error {
	query := `DELETE FROM tournament WHERE id = ?`
	result, err := r.db.ExecContext(ctx, query, tournamentID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("no tournament found to delete")
	}
	return nil
}
