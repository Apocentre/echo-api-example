package server

import (
	"echo-api-example/endpoints"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	e.GET("/", endpoints.HelloWorld)
}
