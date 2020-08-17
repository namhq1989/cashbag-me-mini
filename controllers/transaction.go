package controllers

import (
	"github.com/labstack/echo/v4"

	"cashbag-me-mini/models"
	"cashbag-me-mini/services"
	"cashbag-me-mini/util"
)

// TransactionCreate ...
func TransactionCreate(c echo.Context) error {
	var (
		body = c.Get("body").(models.TransactionCreatePayload)
	)

	// Process data
	rawData, err := services.TransactionCreate(body)

	//if err
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	//Success
	return util.Response200(c, rawData, "")

}


