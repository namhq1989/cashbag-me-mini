package controllers

import (
	"cashbag-me-mini/models"
	"cashbag-me-mini/services"
	"net/http"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

//PatchBranch ...
func PatchBranch(c echo.Context) error {
	id := c.Param("id")
	idBranch, _ := primitive.ObjectIDFromHex(id)
	result := services.PatchBranch(idBranch)
	return c.JSON(http.StatusOK, result)
}

//PutBranch ...
func PutBranch(c echo.Context) error {
	id := c.Param("id")
	idBranch, _ := primitive.ObjectIDFromHex(id)
	body := c.Get("body").(*models.PutBranch)
	result := services.PutBranch(idBranch, *body)
	return c.JSON(http.StatusOK, result)
}
