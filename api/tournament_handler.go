package api

import (
	"VykingProject/models"
	service "VykingProject/service/tournament"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type TournamentAPI struct {
	Service service.TournamentService
}

// CreateTournament godoc
// @Summary Create a new tournament
// @Description Adds a new tournament to the system
// @Tags Tournaments
// @Accept json
// @Produce json
// @Param tournament body models.CreateTournamentRequest true "Tournament data"
// @Success 201 {object} models.Tournament
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Database error"
// @Router /api/tournaments [post]
func (api *TournamentAPI) CreateTournament(w http.ResponseWriter, r *http.Request) {
	var t models.CreateTournamentRequest
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	created, err := api.Service.CreateTournament(r.Context(), t)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
}

// GetAllTournaments godoc
// @Summary Get all tournaments
// @Description Retrieves a list of all tournaments
// @Tags Tournaments
// @Produce json
// @Success 200 {array} models.Tournament
// @Failure 500 {object} map[string]string "Database error"
// @Router /api/tournaments [get]
func (api *TournamentAPI) GetAllTournaments(w http.ResponseWriter, r *http.Request) {
	tournaments, err := api.Service.GetAllTournaments(r.Context())
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(tournaments)
}

// GetTournamentByID godoc
// @Summary Get a tournament by ID
// @Description Retrieves a tournament by its UUID
// @Tags Tournaments
// @Produce json
// @Param id path string true "Tournament ID"
// @Success 200 {object} models.Tournament
// @Failure 400 {object} map[string]string "Invalid ID format"
// @Failure 404 {object} map[string]string "Tournament not found"
// @Router /api/tournaments/{id} [get]
func (api *TournamentAPI) GetTournamentByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	tID, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid tournament ID", http.StatusBadRequest)
		return
	}
	tournament, err := api.Service.GetTournamentByID(r.Context(), tID)
	if err != nil {
		http.Error(w, "Tournament not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(tournament)
}

// DeleteTournament godoc
// @Summary Delete a tournament by ID
// @Description Deletes a tournament from the system
// @Tags Tournaments
// @Param id path string true "Tournament ID"
// @Success 204 "No Content"
// @Failure 400 {object} map[string]string "Invalid ID format"
// @Failure 500 {object} map[string]string "Error deleting tournament"
// @Router /api/tournaments/{id} [delete]
func (api *TournamentAPI) DeleteTournament(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	tID, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid tournament ID", http.StatusBadRequest)
		return
	}
	if err := api.Service.DeleteTournament(r.Context(), tID); err != nil {
		http.Error(w, "Error deleting tournament", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
