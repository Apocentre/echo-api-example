package server

import (
	"echo-api-example/endpoints"
	"fmt"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	fmt.Printf("Inside register routes")
	e.GET("/", endpoints.HelloWorld)
}
