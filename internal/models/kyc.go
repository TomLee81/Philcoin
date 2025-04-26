package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// KYCRecord represents a KYC record in the database
type KYCRecord struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID      string             `bson:"user_id" json:"user_id"`
	Name        string             `bson:"name" json:"name"`
	DOB         string             `bson:"dob" json:"dob"`
	Nationality string             `bson:"nationality" json:"nationality"`
	RawText     string             `bson:"raw_text" json:"raw_text"`
	IPFSCID     string             `bson:"ipfs_cid" json:"ipfs_cid"`
	Status      string             `bson:"status" json:"status"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
}
