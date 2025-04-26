package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Escrow represents an escrow record in the database
type Escrow struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	BuyerID  string             `bson:"buyer_id" json:"buyer_id"`
	SellerID string             `bson:"seller_id" json:"seller_id"`
	Amount   float64            `bson:"amount" json:"amount"`
	Token    string             `bson:"token" json:"token"`
	Status   string             `bson:"status" json:"status"`
}
