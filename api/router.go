package api

import (
	_ "VykingProject/docs"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/gorilla/mux"
)

type RouterConfig struct {
	PlayerAPI     *PlayerAPI
	TournamentAPI *TournamentAPI
	PrizeAPI      *PrizeAPI
}

func NewRouter(cfg RouterConfig) http.Handler {
	r := mux.NewRouter()

	// API Subrouter with /api prefix
	api := r.PathPrefix("/api").Subrouter()

	api.HandleFunc("/players/ranking", cfg.PlayerAPI.GetPlayersByRanking).Methods("GET")
	api.HandleFunc("/players", cfg.PlayerAPI.CreatePlayer).Methods("POST")
	api.HandleFunc("/players", cfg.PlayerAPI.GetAllPlayers).Methods("GET")
	api.HandleFunc("/players/{id}", cfg.PlayerAPI.GetPlayerByID).Methods("GET")
	api.HandleFunc("/players/{id}", cfg.PlayerAPI.DeletePlayer).Methods("DELETE")

	api.HandleFunc("/tournaments", cfg.TournamentAPI.CreateTournament).Methods("POST")
	api.HandleFunc("/tournaments", cfg.TournamentAPI.GetAllTournaments).Methods("GET")
	api.HandleFunc("/tournaments/{id}", cfg.TournamentAPI.GetTournamentByID).Methods("GET")
	api.HandleFunc("/tournaments/{id}", cfg.TournamentAPI.DeleteTournament).Methods("DELETE")

	api.HandleFunc("/prizes/distribute/{id}", cfg.PrizeAPI.DistributePrizes).Methods("PUT")
	api.HandleFunc("/prizes", cfg.PrizeAPI.CreatePrize).Methods("POST")

	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	return r
}
