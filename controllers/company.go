package controllers

import (
	"cashbag-me-mini/models"
	"cashbag-me-mini/services"
	"net/http"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//CreateCompany func to ...
func CreateCompany(c echo.Context) error {
	body := c.Get("body").(*models.PostCompany)
	company := services.CreateCompany(*body)
	return c.JSON(http.StatusCreated, company)
}

//ListCompany to
func ListCompany(c echo.Context) error {
	companies := services.ListCompany()
	return c.JSON(http.StatusOK, companies)
}

//PatchCompany  ...
func PatchCompany(c echo.Context) error {
	id := c.Param("id")
	idCompany, _ := primitive.ObjectIDFromHex(id)
	result := services.PatchBranch(idCompany)
	return c.JSON(http.StatusOK, result)
}
