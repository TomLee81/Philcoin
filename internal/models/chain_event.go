package models

// ChainEvent represents an on-chain event indexed in the database
type ChainEvent struct {
	TxHash    string `bson:"tx_hash" json:"tx_hash"`
	BlockNum  int    `bson:"block_num" json:"block_num"`
	EventName string `bson:"event_name" json:"event_name"`
	Data      []byte `bson:"data" json:"data"`
}
