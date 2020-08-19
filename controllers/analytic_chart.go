package controllers

import (
	"github.com/labstack/echo/v4"

	"cashbag-me-mini/services"
	"cashbag-me-mini/util"
)


// AnalyticChart ...
func AnalyticChart(c echo.Context) error {
	var (
		id          = c.Param("id")
		companyID, _ = util.ValidationObjectID(id)
	)

	// Process data
	rawData, err := services.AnalyticChart(companyID)

	// if err
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	// Success
	return util.Response200(c, rawData, "")
}