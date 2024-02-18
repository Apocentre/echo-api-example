package server

import (
	"echo-api-example/endpoints/eth"
	"echo-api-example/endpoints/hello"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	e.GET("/hello", hello.GetHelloWorld)

	// ETH routes
	eth_group := e.Group("/eth")
	eth_group.GET("/balance", eth.GetAddressBalance)
}
