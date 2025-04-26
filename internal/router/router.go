package router

import (
	"encoding/json"
	"net/http"
	"time"

	"philcoin/internal/handlers"
	"philcoin/internal/middleware"

	"github.com/gorilla/mux"
)

// NewServer returns an *http.Server with all routes and middleware applied
func NewServer(cfg Config, db *mongo.Client) *http.Server {
	r := mux.NewRouter()

	// API versioning
	api := r.PathPrefix("/v1").Subrouter()

	// Apply common middleware
	api.Use(middleware.Logging)
	api.Use(middleware.CORS)
	api.Use(middleware.RateLimit)

	// Register routes
	registerHealthRoutes(api)
	registerMetricsRoutes(api)

	// Default handler for undefined routes
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Endpoint not found"})
	})

	// Server configuration
	return &http.Server{
		Addr:         cfg.ServerAddr,
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
}

// registerHealthRoutes registers health check routes
func registerHealthRoutes(api *mux.Router) {
	api.HandleFunc("/healthz", handlers.Healthz).Methods("GET")
	api.HandleFunc("/readyz", handlers.Readyz).Methods("GET")
}

// registerMetricsRoutes registers metrics routes
func registerMetricsRoutes(api *mux.Router) {
	api.HandleFunc("/metrics", handlers.Metrics).Methods("GET")
}
