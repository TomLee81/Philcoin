package handlers

import (
	"net/http"

	"philcoin/internal/services"
	"philcoin/internal/utils"
)

// GetEvents returns the list of indexed on-chain events
func GetEvents(w http.ResponseWriter, r *http.Request) {
	events, err := services.FetchEvents()
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondJSON(w, http.StatusOK, events)
}
