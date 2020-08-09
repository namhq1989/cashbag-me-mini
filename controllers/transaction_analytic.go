package controllers

import (
	"github.com/labstack/echo/v4"

	"cashbag-me-mini/services"
	"cashbag-me-mini/ultis"
)

// TransactionAnalyticList ...
func TransactionAnalyticList(c echo.Context) error {
	var (
		date = c.QueryParam("date")
	)

	// Process data
	rawData, err := services.TransactionAnalyticList(date)

	// if err
	if err != nil {
		return ultis.Response400(c, nil, err.Error())
	}

	// Success
	return ultis.Response200(c, rawData, "")
}
