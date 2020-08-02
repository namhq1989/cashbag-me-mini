package controllers

import (
	"net/http"
	"cashbag-me-mini/services"
	"github.com/labstack/echo"
)
//CompanyList to 
func CompanyList(c echo.Context) error {
	companies := services.CompanyList()
	return c.JSON(http.StatusOK, companies)
}
