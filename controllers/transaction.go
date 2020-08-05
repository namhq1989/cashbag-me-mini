package controllers

import (
	"cashbag-me-mini/models"
	"cashbag-me-mini/services"
	"net/http"

	"github.com/labstack/echo"
)

//CreateTransaction  ...
func CreateTransaction(c echo.Context) error {
	body := c.Get("body").(*models.PostTransaction)
	result := services.CreateTransaction(*body)
	return c.JSON(http.StatusCreated, result)
}
