package controllers

import (
	"github.com/labstack/echo/v4"

	"cashbag-me-mini/models"
	"cashbag-me-mini/services"
	"cashbag-me-mini/util"
)

// LoyaltyProgramCreate ....
func LoyaltyProgramCreate(c echo.Context) error {
	var (
		body = c.Get("body").(models.LoyaltyProgramCreatePayload)
	)

	// Process data
	rawData, err := services.LoyaltyProgramCreate(body)

	// if err
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	//Success
	return util.Response200(c, rawData, "")
}
