package controllers

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"

	"cashbag-me-mini/models"
	"cashbag-me-mini/services"
	"cashbag-me-mini/util"
)

// BranchList ...
func BranchList(c echo.Context) error {
	// Process data
	rawData, err := services.BranchList()

	// if err
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	// Success
	return util.Response200(c, rawData, "")
}

// BranchCreate ...
func BranchCreate(c echo.Context) error {
	var (
		body = c.Get("body").(models.BranchCreatePayload)
		companyID = c.Get("companyID").(primitive.ObjectID)
	)

	// Process data
	rawData, err := services.BranchCreate(body,companyID)

	// if err
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	// Success
	return util.Response200(c, bson.M{
		"_id":       rawData.ID,
		"createdAt": rawData.CreatedAt,
	}, "")
}

// BranchChangeActiveStatus ...
func BranchChangeActiveStatus(c echo.Context) error {
	var (
		branch = c.Get("branch").(models.BranchBSON)
	)

	// Process data
	rawData, err := services.BranchChangeActiveStatus(branch.ID, !branch.Active)

	// if err
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	// Success
	return util.Response200(c, bson.M{
		"active":    rawData.Active,
		"updatedAt": rawData.UpdatedAt,
	}, "")
}

// BranchUpdate ...
func BranchUpdate(c echo.Context) error {
	var (
		body     = c.Get("body").(models.BranchUpdatePayload)
		branch   = c.Get("branch").(models.BranchBSON)
		branchID = branch.ID
	)

	// Process data
	rawData, err := services.BranchUpdate(branchID, body)

	// if err
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	// Success
	return util.Response200(c, bson.M{
		"updatedAt": rawData.UpdatedAt,
	}, "")
}
