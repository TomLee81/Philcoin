package handlers

import (
	"encoding/json"
	"net/http"
)

// HealthResponse represents the structure of health check responses
type HealthResponse struct {
	Status string `json:"status"`
}

// Healthz checks the liveness of the application
func Healthz(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(HealthResponse{Status: "OK"})
}

// Readyz checks the readiness of the application by verifying dependencies
func Readyz(w http.ResponseWriter, _ *http.Request) {
	// TODO: Implement actual dependency checks (e.g., MongoDB, Redis)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(HealthResponse{Status: "READY"})
}
