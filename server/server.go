package server

import (
	"github.com/labstack/echo/v4"
)

func Start() {
	e := echo.New()
	RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":3000"))
}
