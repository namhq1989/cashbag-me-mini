package controllers

import (
	"cashbag-me-mini/models"
	"cashbag-me-mini/services"
	"net/http"

	"github.com/labstack/echo"
)

//ListBranch ...
func ListBranch(c echo.Context) error {
	Branches := services.ListBranch()
	return c.JSON(http.StatusOK, Branches)
}

//CreateBranch ...
func CreateBranch(c echo.Context) error {
	body := c.Get("body").(*models.PostBranch)
	result := services.CreateBranch(*body)
	return c.JSON(http.StatusOK, result)
}
