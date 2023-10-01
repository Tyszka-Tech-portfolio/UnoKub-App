package utils

import (
	"github.com/ethereum/go-ethereum/rpc"
)

// InitializeWeb3Client initializes a new Web3 client
func InitializeWeb3Client(url string) (*rpc.Client, error) {
	client, err := rpc.Dial(url)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// TODO: Add more Web3 related helper functions
// e.g., SendTransaction, QuerySmartContract, etc.
