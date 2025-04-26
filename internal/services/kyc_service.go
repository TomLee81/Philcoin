package services

import (
	"io"
	"time"

	"philcoin/internal/models"
	"philcoin/internal/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ProcessIDImage performs OCR on the uploaded ID image, parses fields, and saves a KYCRecord
func ProcessIDImage(reader io.Reader) (*models.KYCRecord, error) {
	// Upload to IPFS
	cid, err := utils.UploadToIPFS(reader)
	if err != nil {
		return nil, err
	}

	// OCR
	text, err := utils.OCRExtractText(reader)
	if err != nil {
		return nil, err
	}

	data := utils.ParseOCR(text)

	rec := &models.KYCRecord{
		ID:          primitive.NewObjectID(),
		UserID:      data.UserID,
		Name:        data.Name,
		DOB:         data.DOB,
		Nationality: data.Nationality,
		RawText:     text,
		IPFSCID:     cid,
		Status:      "pending",
		CreatedAt:   time.Now(),
	}
	if err := SaveKYCRecord(rec); err != nil {
		return nil, err
	}
	return rec, nil
}

// VerifyLiveness calls an external AI service to check liveness
func VerifyLiveness(reader io.Reader) (bool, error) {
	return utils.CallLivenessAPI(reader)
}

// SaveKYCRecord saves the KYC record to the database
func SaveKYCRecord(rec *models.KYCRecord) error {
	// Implement database save logic here
	return nil
}
