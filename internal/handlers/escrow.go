package handlers

import (
	"encoding/json"
	"net/http"

	"philcoin/internal/services"
	"philcoin/internal/utils"

	"github.com/gorilla/mux"
)

// CreateEscrow initializes a new escrow agreement
func CreateEscrow(w http.ResponseWriter, r *http.Request) {
	var req services.EscrowRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Validation
	if req.Amount <= 0 {
		utils.RespondError(w, http.StatusBadRequest, "Amount must be positive")
		return
	}
	if !services.IsSupportedToken(req.Token) {
		utils.RespondError(w, http.StatusBadRequest, "Unsupported token")
		return
	}

	esc, err := services.CreateEscrow(req)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondJSON(w, http.StatusCreated, esc)
}

// ReleaseEscrow releases funds from an existing escrow
func ReleaseEscrow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["escrowId"]

	if err := services.ReleaseEscrow(id); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	utils.RespondJSON(w, http.StatusOK, map[string]string{"status": "released"})
}
