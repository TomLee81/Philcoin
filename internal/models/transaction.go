package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Transaction represents a blockchain transaction
type Transaction struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	FromWallet primitive.ObjectID `bson:"from_wallet" json:"fromWallet"`
	ToWallet   primitive.ObjectID `bson:"to_wallet" json:"toWallet"`
	Amount     float64            `bson:"amount" json:"amount"`
	Token      string             `bson:"token" json:"token"`
	TxHash     string             `bson:"tx_hash" json:"txHash"`
	Status     string             `bson:"status" json:"status"`
	CreatedAt  time.Time          `bson:"created_at" json:"createdAt"`
}
