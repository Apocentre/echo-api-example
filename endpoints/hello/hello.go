package hello

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetHelloWorld(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "hello, world!")
}
