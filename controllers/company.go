package controllers

import (
	"cashbag-me-mini/services"
	"net/http"

	"github.com/labstack/echo"
)

//CompanyList to
func CompanyList(c echo.Context) error {
	companies := services.CompanyList()
	return c.JSON(http.StatusOK, companies)
}

//CompanyCreate func to ...
func CompanyCreate(c echo.Context) error {
	company := services.CompanyCreate(c)
	return c.JSON(http.StatusCreated, company)
}
//CompanyUpdate func to ... 
func CompanyUpdate(c echo.Context) error{
	company :=services.CompanyUpdate(c)
	return c.JSON(http.StatusOK,company)
}

//CompanyActive func to ... 
func CompanyActive(c echo.Context) error{
	company :=services.CompanyActive(c)
	return c.JSON(http.StatusOK,company)
}
