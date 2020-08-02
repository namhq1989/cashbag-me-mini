package controllers

import (
	"cashbag-me-mini/services"
	"net/http"

	"github.com/labstack/echo"
)

//ListBranch ...
func ListBranch(c echo.Context) error {
	Branchs := services.ListBranch()
	return c.JSON(http.StatusOK, Branchs)
}

