package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"philcoin/internal/handlers"

	"github.com/stretchr/testify/assert"
)

func TestMetricsHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/metrics", nil)
	rr := httptest.NewRecorder()
	handlers.Metrics(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), "# HELP")
}
