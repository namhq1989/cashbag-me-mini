package controllers

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"

	"cashbag-me-mini/models"
	"cashbag-me-mini/services"
	"cashbag-me-mini/util"
)

// CompanyCreate ...
func CompanyCreate(c echo.Context) error {
	var (
		body = c.Get("body").(models.CompanyCreatePayload)
	)

	// Process data
	rawData, err := services.CompanyCreate(body)

	// if err
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	//Success
	return util.Response200(c, bson.M{
		"_id":       rawData.ID,
		"createdAt": rawData.CreatedAt,
	}, "")
}

// CompanyList ...
func CompanyList(c echo.Context) error {
	// Process data
	rawData, err := services.CompanyList()

	// if err
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	//success
	return util.Response200(c, rawData, "")
}

// CompanyChangeActiveStatus ...
func CompanyChangeActiveStatus(c echo.Context) error {
	var (
		company = c.Get("company").(models.CompanyBSON)
	)

	// Process data
	rawData, err := services.CompanyChangeActiveStatus(company.ID, !company.Active)

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

// CompanyUpdate ...
func CompanyUpdate(c echo.Context) error {
	var (
		body      = c.Get("body").(models.CompanyUpdatePayload)
		company   = c.Get("company").(models.CompanyBSON)
		companyID = company.ID
	)

	// Process data
	rawData, err := services.CompanyUpdate(companyID, body)

	// if err
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	// Success
	return util.Response200(c, bson.M{
		"updatedAt": rawData.UpdatedAt,
	}, "")
}
