package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"philcoin/internal/handlers"
	"philcoin/internal/models"
	"philcoin/internal/services"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestCreateEscrow_Success(t *testing.T) {
	// Mock CreateEscrow
	orig := services.CreateEscrow
	services.CreateEscrow = func(req services.EscrowRequest) (*models.Escrow, error) {
		return &models.Escrow{
			ID:       primitive.NewObjectID(),
			BuyerID:  req.Buyer,
			SellerID: req.Seller,
			Amount:   req.Amount,
			Token:    req.Token,
			Status:   "locked",
		}, nil
	}
	defer func() { services.CreateEscrow = orig }()

	payload := map[string]interface{}{
		"buyer":  "A",
		"seller": "B",
		"amount": 100.0,
		"token":  "PHIL",
	}
	body, _ := json.Marshal(payload)
	req := httptest.NewRequest("POST", "/v1/escrow", bytes.NewReader(body))
	rr := httptest.NewRecorder()

	handlers.CreateEscrow(rr, req)
	assert.Equal(t, http.StatusCreated, rr.Code)
}

func TestReleaseEscrow_Success(t *testing.T) {
	req := httptest.NewRequest("POST", "/v1/escrow/123/release", nil)
	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/v1/escrow/{escrowId}/release", handlers.ReleaseEscrow)
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}
