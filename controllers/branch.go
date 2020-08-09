package controllers

import (
	"github.com/labstack/echo/v4"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"cashbag-me-mini/models"
	"cashbag-me-mini/services"
	"cashbag-me-mini/ultis"
)

// BranchList ...
func BranchList(c echo.Context) error {
	// Process data
	rawData, err := services.BranchList()

	// if err
	if err != nil {
		return ultis.Response400(c, nil, err.Error())
	}

	// Success
	return ultis.Response200(c, rawData, "")
}

// BranchCreate ...
func BranchCreate(c echo.Context) error {
	var (
		body = c.Get("body").(*models.BranchCreatePayload)
	)

	// Process data
	rawData, err := services.BranchCreate(*body)

	// if err
	if err != nil {
		return ultis.Response400(c, nil, err.Error())
	}

	// Success
	return ultis.Response200(c, rawData, "")
}

// BranchChangeActiveStatus ...
func BranchChangeActiveStatus(c echo.Context) error {
	var (
		id          = c.Param("id")
		branchID, _ = primitive.ObjectIDFromHex(id)
	)

	// Process data
	rawData, err := services.BranchChangeActiveStatus(branchID)

	// if err
	if err != nil {
		return ultis.Response400(c, nil, err.Error())
	}

	// Success
	return ultis.Response200(c, rawData, "")
}

// BranchUpdate ...
func BranchUpdate(c echo.Context) error {
	var (
		id          = c.Param("id")
		body        = c.Get("body").(*models.BranchUpdateBPayload)
		branchID, _ = primitive.ObjectIDFromHex(id)
	)

	// Process data
	rawData, err := services.BranchUpdate(branchID, *body)

	// if err
	if err != nil {
		return ultis.Response400(c, nil, err.Error())
	}

	// Success
	return ultis.Response200(c, rawData, "")
}
