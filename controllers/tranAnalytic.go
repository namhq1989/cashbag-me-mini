package controllers

import (
	"cashbag-me-mini/services"
	"net/http"

	"github.com/labstack/echo"
)

//TranAnalytic ...
func TranAnalytic(c echo.Context) error {
	date := c.QueryParam("date")
	tranAnalytic := services.TranAnalytic(date)
	return c.JSON(http.StatusOK, tranAnalytic)
}
