package models

import (
	"errors"
	"time"
)

// NFTType represents the type of NFT
type NFTType int

const (
	PlayerCard NFTType = iota
	OtherAsset
)

// NFT struct represents a non-fungible token
type NFT struct {
	TokenID     string    `json:"token_id"`
	UserID      string    `json:"user_id"`
	OwnerID     string    `json:"owner_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Metadata    string    `json:"metadata"`
	Type        NFTType   `json:"type"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// NewNFT initializes a new NFT
func NewNFT(tokenID, userID, ownerID, name, description, metadata string, nftType NFTType, existingNFTs []NFT) (*NFT, error) {
	// Check if user already has a player card
	if nftType == PlayerCard {
		for _, nft := range existingNFTs {
			if nft.UserID == userID && nft.Type == PlayerCard {
				return nil, errors.New("User already has a player card NFT")
			}
		}
	}

	return &NFT{
		TokenID:     tokenID,
		UserID:      userID,
		OwnerID:     ownerID,
		Name:        name,
		Description: description,
		Metadata:    metadata,
		Type:        nftType,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}

// Other existing methods...
