package server

import (
	"echo-api-example/blockchain"

	"github.com/labstack/echo/v4"
)

/// Custom context that extends the default Echo context
type ApiContext struct {
	echo.Context
	eth_client * blockchain.EthClient
}

func Start(eth_client * blockchain.EthClient) {
	e := echo.New()

	// middlewares

	// Extend the default echo context
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &ApiContext{c, eth_client}
			return next(cc)
		}
	})

	RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":3000"))
}
