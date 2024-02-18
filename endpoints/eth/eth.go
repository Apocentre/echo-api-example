package eth

import (
	"fmt"
	"net/http"

	"echo-api-example/endpoints"

	"github.com/labstack/echo/v4"
)

type QS struct {
  Address string `query:"address"`
}

func GetAddressBalance(c echo.Context) error {
	fmt.Println("Running...")
	ctx := c.(*endpoints.ApiContext)

	var qs QS
	err := c.Bind(&qs); if err != nil {
    return c.String(http.StatusBadRequest, "bad request")
	}

	balance, err := ctx.EthClient().GetBalance(qs.Address)
	if err != nil {
		return err
	}

	return ctx.String(http.StatusOK, balance.String())
}
