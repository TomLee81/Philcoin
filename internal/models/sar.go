package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SARReport represents a Suspicious Activity Report
type SARReport struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID      primitive.ObjectID `bson:"user_id" json:"userId"`
	Description string             `bson:"description" json:"description"`
	CreatedAt   time.Time          `bson:"created_at" json:"createdAt"`
}
