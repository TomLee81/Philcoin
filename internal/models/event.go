package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ChainEvent indexes on-chain smart contract events
type ChainEvent struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	TxHash    string             `bson:"tx_hash" json:"txHash"`
	BlockNum  uint64             `bson:"block_num" json:"blockNum"`
	EventName string             `bson:"event_name" json:"eventName"`
	Data      []byte             `bson:"data" json:"data"`
	Timestamp time.Time          `bson:"timestamp" json:"timestamp"`
}
