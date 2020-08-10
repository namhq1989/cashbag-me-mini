package controllers

import (
	"github.com/labstack/echo/v4"

	"cashbag-me-mini/models"
	"cashbag-me-mini/services"
	"cashbag-me-mini/ultis"
)

// TransactionCreate ...
func TransactionCreate(c echo.Context) error {
	var (
		body = c.Get("body").(*models.TransactionCreatePayload)
	)

	// Process data
	rawData, err := services.TransactionCreate(*body)

	//if err
	if err != nil {
		return ultis.Response400(c, nil, err.Error())
	}

	//Success
	return ultis.Response200(c, rawData, "")

}
