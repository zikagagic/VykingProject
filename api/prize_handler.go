package api

import (
	"VykingProject/models"
	service "VykingProject/service/prize"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type PrizeAPI struct {
	Service service.PrizeService
}

// CreatePrize godoc
// @Summary Create prize that references the tournament and the player's placement
// @Description Creates prize with placement in mind
// @Tags Prizes
// @Param player body models.CreatePrizeRequest true "Player data"
// @Success 204 {object} models.Prize
// @Failure 500 {object} map[string]string "Error creating prize"
// @Router /api/prizes [post]
func (api *PrizeAPI) CreatePrize(w http.ResponseWriter, r *http.Request) {
	var request models.CreatePrizeRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid input for creating prize", http.StatusBadRequest)
	}
	if err := api.Service.CreatePrize(r.Context(), request); err != nil {
		http.Error(w, "Error writing prize to database", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(request)
}

// DistributePrizes godoc
// @Summary Distribute tournament prizes
// @Description Distributes tournament prizes to the top 3 people placed
// @Tags Prizes
// @Param id path string true "Tournament ID"
// @Success 204 "No Content"
// @Failure 500 {object} map[string]string "Error distributing prizes"
// @Router /api/prizes/distribute/{id} [put]
func (api *PrizeAPI) DistributePrizes(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	tournamentID, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid tournament ID", http.StatusBadRequest)
		return
	}
	err = api.Service.DistributePrizes(r.Context(), tournamentID)
	if err != nil {
		http.Error(w, "Failed to get prize report", http.StatusBadRequest)
		return
	}
}
