package repository

import (
	"context"
	"database/sql"
	"errors"

	"VykingProject/models"

	"github.com/google/uuid"
)

type MYSQLPlayerRepository struct {
	db *sql.DB
}

func NewMYSQLPlayerRepository(db *sql.DB) *MYSQLPlayerRepository {
	return &MYSQLPlayerRepository{db: db}
}

// Inserts a new player record into the database.
func (r *MYSQLPlayerRepository) CreatePlayer(ctx context.Context, player models.Player) error {
	query := `INSERT INTO player (id, name, email, account_balance) VALUES (?, ?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, player.ID, player.Name, player.Email, player.AccountBalance)
	return err
}

// Deletes a player record by ID.
func (r *MYSQLPlayerRepository) DeletePlayer(ctx context.Context, playerID uuid.UUID) error {
	query := `DELETE FROM player WHERE id = ?`
	result, err := r.db.ExecContext(ctx, query, playerID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("no player found to delete")
	}
	return nil
}

// Retrieves all players from the database.
func (r *MYSQLPlayerRepository) GetAllPlayers(ctx context.Context) ([]models.Player, error) {
	query := `SELECT id, name, email, account_balance FROM player`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var players []models.Player
	for rows.Next() {
		var p models.Player
		err := rows.Scan(&p.ID, &p.Name, &p.Email, &p.AccountBalance)
		if err != nil {
			return nil, err
		}
		players = append(players, p)
	}
	return players, rows.Err()
}

func (r *MYSQLPlayerRepository) GetPlayersByRanking(ctx context.Context) ([]models.RankedPlayer, error) {
	// Window Function for returning ranked players
	query := "WITH RankedPlayers AS ( " +
		"SELECT id, name, email, account_balance, RANK() OVER (ORDER BY account_balance DESC) AS `rank` " +
		"FROM player ) " +
		"SELECT id, name, email, account_balance, `rank` FROM RankedPlayers;"

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var players []models.RankedPlayer
	for rows.Next() {
		var p models.RankedPlayer
		err := rows.Scan(&p.ID, &p.Name, &p.Email, &p.AccountBalance, &p.Rank)
		if err != nil {
			return nil, err
		}
		players = append(players, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return players, nil
}

// Fetches a player by ID.
func (r *MYSQLPlayerRepository) GetPlayerByID(ctx context.Context, playerID uuid.UUID) (models.Player, error) {
	query := `SELECT id, name, email, account_balance FROM player WHERE id = ?`
	var p models.Player
	err := r.db.QueryRowContext(ctx, query, playerID).Scan(&p.ID, &p.Name, &p.Email, &p.AccountBalance)
	if errors.Is(err, sql.ErrNoRows) {
		return p, errors.New("player not found")
	}
	return p, err
}
