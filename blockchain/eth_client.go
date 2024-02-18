package blockchain

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
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

func (c *EthClient) GetBalance(hex_address string) (*big.Int, error) {
	address := common.HexToAddress(hex_address)
	balance, err := c.client.BalanceAt(context.Background(), address, nil)

	if err != nil {
		return nil, fmt.Errorf("error reading balance of %s : %v", hex_address, err)
	}

	return balance, nil
}
