package eth

import (
	"fmt"
	"math/big"
	"net/http"
	"time"

	"echo-api-example/endpoints"

	"github.com/labstack/echo/v4"
)

type QS struct {
  Address string `query:"address"`
}

type RequestError struct {
	Status int
	Msg string
}

func NewRequestError(status int, msg string) *RequestError {
	return &RequestError {status, msg}
}

func GetAddressBalance(c echo.Context) error {
	fmt.Println("Running...")
	result_c := make(chan *big.Int, 1)
	error_c := make(chan *RequestError, 1)
	ctx := c.(*endpoints.ApiContext)
	eth_client := ctx.EthClient()
	
	go func ()  {
		time.Sleep(10 * time.Second)
	
		var qs QS
		err := c.Bind(&qs); if err != nil {
			error_c <- NewRequestError(http.StatusBadRequest, "bad request")
		}
	
		balance, err := eth_client.GetBalance(qs.Address)
		if err != nil {
			error_c <- NewRequestError(http.StatusInternalServerError, "internal server error")
		}	

		result_c <- balance
	}()
	
	select {
	case balance := <- result_c:
		return ctx.String(http.StatusOK, balance.String())
	case req_error := <- error_c:
		return c.String(req_error.Status, req_error.Msg)
	case <- c.Request().Context().Done():
		return nil
	}
	
}
