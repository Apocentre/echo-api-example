package endpoints

import (
	"echo-api-example/blockchain"

	"github.com/labstack/echo/v4"
)

/// Custom context that extends the default Echo context
type ApiContext struct {
	echo.Context
	eth_client * blockchain.EthClient
}

func NewApiContext(c echo.Context, eth_client *blockchain.EthClient) *ApiContext {
	return &ApiContext {c, eth_client}
}

func (c *ApiContext) EthClient() * blockchain.EthClient {
	return c.eth_client
}
