package api

import (
	"VykingProject/models"
	service "VykingProject/service/player"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type PlayerAPI struct {
	Service service.PlayerService
}

// CreatePlayer godoc
// @Summary Create a new player
// @Description Adds a new player to the system
// @Tags Players
// @Accept json
// @Produce json
// @Param player body models.CreatePlayerRequest true "Player data"
// @Success 201 {object} models.CreatePlayerRequest
// @Failure 400 {object} map[string]string
// @Router /api/players [post]
func (api *PlayerAPI) CreatePlayer(w http.ResponseWriter, r *http.Request) {
	var request models.CreatePlayerRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	createdPlayer, err := api.Service.CreatePlayer(r.Context(), request)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdPlayer)
}

// GetAllPlayers godoc
// @Summary Get all players
// @Description Retrieves a list of all players
// @Tags Players
// @Produce json
// @Success 200 {array} models.Player
// @Failure 500 {object} map[string]string
// @Router /api/players [get]
func (api *PlayerAPI) GetAllPlayers(w http.ResponseWriter, r *http.Request) {
	players, err := api.Service.GetAllPlayers(r.Context())
	if err != nil {
		http.Error(w, "Failed to retrieve players", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(players)
}

// GetPlayerByID godoc
// @Summary Get a player by ID
// @Description Retrieves a player by their unique ID
// @Tags Players
// @Produce json
// @Param id path string true "Player ID"
// @Success 200 {object} models.Player
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/players/{id} [get]
func (api *PlayerAPI) GetPlayerByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	playerID, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid player ID", http.StatusBadRequest)
		return
	}
	player, err := api.Service.GetPlayerByID(r.Context(), playerID)
	if err != nil {
		http.Error(w, "Player not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(player)
}

// GetPlayersByRanking godoc
// @Summary Get players sorted by ranking
// @Description Retrieves players ordered by tournament ranking
// @Tags Players
// @Produce json
// @Success 200 {array} models.RankedPlayer
// @Failure 500 {object} map[string]string
// @Router /api/players/ranking [get]
func (api *PlayerAPI) GetPlayersByRanking(w http.ResponseWriter, r *http.Request) {
	players, err := api.Service.GetPlayersByRanking(r.Context())
	if err != nil {
		http.Error(w, "Could not retrieve players by ranking", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(players)
}

// DeletePlayer godoc
// @Summary Delete a player
// @Description Deletes a player by their unique ID
// @Tags Players
// @Param id path string true "Player ID"
// @Success 204 "No Content"
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/players/{id} [delete]
func (api *PlayerAPI) DeletePlayer(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	playerID, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid player ID", http.StatusBadRequest)
		return
	}
	if err := api.Service.DeletePlayer(r.Context(), playerID); err != nil {
		http.Error(w, "Error deleting player", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
