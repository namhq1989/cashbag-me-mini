package controllers

import (
	"cashbag-me-mini/services"
	"cashbag-me-mini/models"
	"net/http"

	"github.com/labstack/echo"
)
//CreateCompany func to ...
func CreateCompany(c echo.Context) error {
	body := c.Get("body").(*models.PostCompany)
	company := services.CreateCompany(*body)
	return c.JSON(http.StatusCreated, company)
}


