package controllers

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"cashbag-me-mini/models"
	"cashbag-me-mini/services"
	"cashbag-me-mini/ultis"
)

// UserCreate ...
func UserCreate(c echo.Context) error {
	var (
		body = c.Get("body").(*models.UserCreatePayload)
	)

	//Process data
	rawData, err := services.UserCreate(*body)
	if err != nil {
		return ultis.Response400(c, nil, err.Error())
	}

	//Success
	return ultis.Response200(c, rawData, "")

}

// UserUpdate ...
func UserUpdate(c echo.Context) error{
	var (
		id           = c.Param("id")
		body         = c.Get("body").(*models.UserUpdatePayload)
		userID, _ = primitive.ObjectIDFromHex(id)
	)

	// Process data
	rawData, err := services.UserUpdate(userID, *body)

	// if err
	if err != nil {
		return ultis.Response400(c, nil, err.Error())
	}

	//success
	return ultis.Response200(c, rawData, "")
}

