package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

// Connect establishes a MongoDB connection
func Connect(uri string) *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("MongoDB connection error: %v", err)
	}

	log.Println("✅ Connected to MongoDB")
	return client
}

// Disconnect closes the MongoDB connection
func Disconnect(ctx context.Context) {
	if err := client.Disconnect(ctx); err != nil {
		log.Printf("MongoDB disconnect error: %v", err)
	} else {
		log.Println("✅ Disconnected from MongoDB")
	}
}
