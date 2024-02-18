package server

import (
	"echo-api-example/blockchain"
	"echo-api-example/endpoints"

	"github.com/labstack/echo/v4"
)

func Start(eth_client * blockchain.EthClient) {
	e := echo.New()

	// middlewares

	// Extend the default echo context
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := endpoints.NewApiContext(c, eth_client)
			return next(cc)
		}
	})

	RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":3000"))
}
