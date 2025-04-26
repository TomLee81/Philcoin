package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Wallet represents a user wallet
type Wallet struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	OwnerID   primitive.ObjectID `bson:"owner_id" json:"ownerId"`
	Balance   float64            `bson:"balance" json:"balance"`
	CreatedAt time.Time          `bson:"created_at" json:"createdAt"`
}
