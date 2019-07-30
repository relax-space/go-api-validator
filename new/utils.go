package main

import (
	"fmt"

	"nomni/utils/api"

	"github.com/labstack/echo"
)

func ReturnApiFail(c echo.Context, status int, apiError api.Error, detail ...interface{}) error {

	for _, d := range detail {
		if d != nil {
			apiError.Details = fmt.Sprint(detail...)
		}
	}
	return c.JSON(status, api.Result{
		Success: false,
		Error:   apiError,
	})
}

func ReturnApiSucc(c echo.Context, status int, result interface{}) error {
	return c.JSON(status, api.Result{
		Success: true,
		Result:  result,
	})
}
