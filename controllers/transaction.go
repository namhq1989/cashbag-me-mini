package controllers

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"cashbag-me-mini/models"
	"cashbag-me-mini/services"
	"cashbag-me-mini/util"
)

// TransactionCreate ...
func TransactionCreate(c echo.Context) error {
	var (
		body    = c.Get("body").(models.TransactionCreatePayload)
		company = c.Get("company").(models.CompanyBSON)
		branch  = c.Get("branch").(models.BranchBSON)
		user    = c.Get("user").(models.UserBSON)
	)

	// Process data
	rawData, err := services.TransactionCreate(body, company, branch, user)

	//if err
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	return util.Response200(c, bson.M{
		"_id":       rawData.ID,
		"createdAt": rawData.CreatedAt,
	}, "")

}

// TransactionFindByUserID ...
func TransactionFindByUserID(c echo.Context) error {
	var (
		userID = c.Get("userID").(primitive.ObjectID)
	)

	// process data
	rawData, err := services.TransactionFindByUserID(userID)

	// if err
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	// Success
	return util.Response200(c, rawData, "")
}
