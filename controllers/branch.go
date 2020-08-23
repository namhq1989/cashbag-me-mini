package controllers

import (
	"time"

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
	)

	// Process data
	rawData, err := services.BranchCreate(body)

	// if err
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	// Success
	return util.Response200(c, bson.M{
		"_id":rawData.ID,
		"createdAt":time.Now(),
	}, "")
}

// BranchChangeActiveStatus ...
func BranchChangeActiveStatus(c echo.Context) error {
	var (
		branch = c.Get("branch").(models.BranchBSON)
	)

	// Process data
	_, err := services.BranchChangeActiveStatus(branch.ID, !branch.Active)

	// if err
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	// Success
	return util.Response200(c, bson.M{
		"active": !branch.Active,
		"updatedAt":time.Now(),
	}, "")
}

// BranchUpdate ...
func BranchUpdate(c echo.Context) error {
	var (
		id          = c.Param("id")
		body        = c.Get("body").(models.BranchUpdatePayload)
		branchID, _ = util.ValidationObjectID(id)
	)

	// Process data
	_, err := services.BranchUpdate(branchID, body)

	// if err
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	// Success
	return util.Response200(c, bson.M{
		"updatedAt": time.Now(),
	}, "")
}
