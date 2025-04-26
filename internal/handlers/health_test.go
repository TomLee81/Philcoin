package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"philcoin/internal/handlers"

	"github.com/stretchr/testify/assert"
)

func TestHealthz(t *testing.T) {
	req := httptest.NewRequest("GET", "/healthz", nil)
	rr := httptest.NewRecorder()
	handlers.Healthz(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var resp struct{ Status string }
	json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.Equal(t, "OK", resp.Status)
}

func TestReadyz(t *testing.T) {
	req := httptest.NewRequest("GET", "/readyz", nil)
	rr := httptest.NewRecorder()
	handlers.Readyz(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var resp struct{ Status string }
	json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.Equal(t, "READY", resp.Status)
}
