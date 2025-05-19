// @title Vyking Project API
// @version 1.0
// @description Test project for an iGaming platform
// @host localhost:8080
// @BasePath /

package main

import (
	"VykingProject/api"
	config "VykingProject/config"
	db "VykingProject/db"
	"database/sql"
	"time"

	playerRepo "VykingProject/repository/player"
	prizeRepo "VykingProject/repository/prize"
	tournamentRepo "VykingProject/repository/tournament"

	playerService "VykingProject/service/player"
	prizeService "VykingProject/service/prize"
	tournamentService "VykingProject/service/tournament"

	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func waitForDB() (*sql.DB, error) {
	var db *sql.DB
	var err error

	maxRetries := 10
	for i := 0; i < maxRetries; i++ {
		db, err = config.GetMYSQLDB()
		if err == nil {
			err = db.Ping()
			if err == nil {
				return db, nil
			}
		}
		log.Printf("Waiting for DB to be ready (attempt %d/%d): %v", i+1, maxRetries, err)
		time.Sleep(3 * time.Second)
	}

	return nil, err
}

func main() {
	// Open DB connection
	database, err := waitForDB()
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	log.Println("Connection established to DB")
	defer database.Close()

	// Run migrations
	if err := db.InitMigrations(database); err != nil {
		log.Fatalf("Migration error: %v", err)
	}

	// Repositories
	playerRepo := playerRepo.NewMYSQLPlayerRepository(database)
	tournamentRepo := tournamentRepo.NewMYSQLTournamentRepository(database)
	prizeRepo := prizeRepo.NewMYSQLTournamentRepository(database)

	// Services
	playerService := playerService.NewPlayerService(playerRepo)
	tournamentService := tournamentService.NewTournamentService(tournamentRepo)
	prizeService := prizeService.NewPrizeService(prizeRepo)

	// APIs
	playerAPI := &api.PlayerAPI{Service: playerService}
	tournamentAPI := &api.TournamentAPI{Service: tournamentService}
	prizeAPI := &api.PrizeAPI{Service: prizeService}

	// Router
	router := api.NewRouter(api.RouterConfig{
		PlayerAPI:     playerAPI,
		TournamentAPI: tournamentAPI,
		PrizeAPI:      prizeAPI,
	})

	// Start server
	log.Println("Server started on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
