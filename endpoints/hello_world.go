package endpoints

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HelloWorld(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "hello, world!")
}
