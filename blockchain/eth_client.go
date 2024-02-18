package blockchain

import (
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
)

type EthClient struct {
	client *ethclient.Client
}


func NewEthClient(rpc string) (*EthClient, error) {
	client, err := ethclient.Dial(rpc)
	if err != nil {
		return nil, fmt.Errorf("error connecting to RPC endpoint: %v", err)
	}

	return &EthClient{
		client,
	}, nil
}
