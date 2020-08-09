package controllers

import (
	"cashbag-me-mini/services"
	"net/http"

	"github.com/labstack/echo"
)

//TranAnalytic ...
func TransactionAnalytic(c echo.Context) error {
	var(
		date = c.QueryParam("date")
	)
	
	// Process data
	rawData, err := services.TransactionAnalytic(date)

	// if err
	if err != nil {
		return ultis.Response400(c, nil, err.Error())
	}

	// Success
	return ultis.Response200(c, rawData, "")
}
