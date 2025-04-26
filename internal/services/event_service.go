package services

import (
	"context"

	"philcoin/internal/models"

	"go.mongodb.org/mongo-driver/bson"
)

// FetchEvents retrieves all indexed ChainEvents from the database
func FetchEvents() ([]*models.ChainEvent, error) {
	coll := getDB().Database("philcoin").Collection("chain_events")
	cursor, err := coll.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var events []*models.ChainEvent
	for cursor.Next(context.Background()) {
		var evt models.ChainEvent
		if err := cursor.Decode(&evt); err != nil {
			return nil, err
		}
		events = append(events, &evt)
	}
	return events, nil
}

// StartEventIndexer runs the background event indexer
func StartEventIndexer(rpcURL, contractAddr string) {
	// TODO: Implement the event indexing logic
}
