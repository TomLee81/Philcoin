package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"philcoin/internal/handlers"
	"philcoin/internal/models"
	"philcoin/internal/services"

	"github.com/stretchr/testify/assert"
)

func TestGetEvents_Success(t *testing.T) {
	// Mock FetchEvents
	orig := services.FetchEvents
	services.FetchEvents = func() ([]*models.ChainEvent, error) {
		return []*models.ChainEvent{{
			TxHash:    "0x123",
			BlockNum:  100,
			EventName: "TestEvent",
			Data:      []byte("data"),
		}}, nil
	}
	defer func() { services.FetchEvents = orig }()

	req := httptest.NewRequest("GET", "/v1/events", nil)
	rr := httptest.NewRecorder()

	handlers.GetEvents(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)

	var events []*models.ChainEvent
	err := json.Unmarshal(rr.Body.Bytes(), &events)
	assert.NoError(t, err)
	assert.Len(t, events, 1)
	assert.Equal(t, "TestEvent", events[0].EventName)
}
