package controllers

import (
	"github.com/labstack/echo/v4"

	"go.mongodb.org/mongo-driver/bson/primitive"

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
	return util.Response200(c, rawData, "")
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
		id           = c.Param("id")
		companyID, _ = primitive.ObjectIDFromHex(id)
	)

	// Process data
	rawData, err := services.CompanyChangeActiveStatus(companyID)

	// if err
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	// Success
	return util.Response200(c, rawData, "")
}

// CompanyUpdate ...
func CompanyUpdate(c echo.Context) error {
	var (
		id           = c.Param("id")
		body         = c.Get("body").(models.CompanyUpdatePayload)
		companyID, _ = primitive.ObjectIDFromHex(id)
	)

	// Process data
	rawData, err := services.CompanyUpdate(companyID, body)

	// if err
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	//success
	return util.Response200(c, rawData, "")
}
