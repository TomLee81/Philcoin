package services

import (
	"errors"

	"philcoin/internal/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// EscrowRequest defines the input for creating an escrow
type EscrowRequest struct {
	Buyer  string  `json:"buyer"`
	Seller string  `json:"seller"`
	Amount float64 `json:"amount"`
	Token  string  `json:"token"`
}

// CreateEscrow creates and stores an escrow record
func CreateEscrow(req EscrowRequest) (*models.Escrow, error) {
	esc := &models.Escrow{
		ID:       primitive.NewObjectID(),
		BuyerID:  req.Buyer,
		SellerID: req.Seller,
		Amount:   req.Amount,
		Token:    req.Token,
		Status:   "locked",
	}
	if err := SaveEscrow(esc); err != nil {
		return nil, err
	}
	// TODO: integrate on-chain deposit
	return esc, nil
}

// ReleaseEscrow transitions an escrow to released state
func ReleaseEscrow(id string) error {
	esc, err := FindEscrowByID(id)
	if err != nil {
		return err
	}
	if esc.Status != "locked" {
		return errors.New("escrow not in locked state")
	}
	// TODO: integrate on-chain release
	esc.Status = "released"
	return UpdateEscrow(esc)
}

// IsSupportedToken checks if the token is supported
func IsSupportedToken(token string) bool {
	supportedTokens := map[string]bool{
		"PHIL": true,
		"ETH":  true,
		"BTC":  true,
	}
	return supportedTokens[token]
}
