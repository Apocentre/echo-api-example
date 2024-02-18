package eth

import (
	"net/http"

	"echo-api-example/endpoints"

	"github.com/labstack/echo/v4"
)

type QS struct {
  address string `query:"address"`
}

func GetAddressBalance(c echo.Context) error {
	ctx := c.(*endpoints.ApiContext)

	var qs QS
	err := c.Bind(&qs); if err != nil {
    return c.String(http.StatusBadRequest, "bad request")
	}

	balance, err := ctx.EthClient().GetBalance(qs.address)
	if err != nil {
		return err
	}

	return ctx.String(http.StatusOK, balance.String())
}
