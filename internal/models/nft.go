package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// NFT represents an on-chain NFT asset
type NFT struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	OwnerID     primitive.ObjectID `bson:"owner_id" json:"ownerId"`
	TokenID     string             `bson:"token_id" json:"tokenId"`
	MetadataURI string             `bson:"metadata_uri" json:"metadataUri"`
	CreatedAt   time.Time          `bson:"created_at" json:"createdAt"`
}
