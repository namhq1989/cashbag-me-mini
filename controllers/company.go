package controllers

import (
	"cashbag-me-mini/models"
	"cashbag-me-mini/services"
	"cashbag-me-mini/ultis"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CompanyCreate func to ...
func CompanyCreate(c echo.Context) error {
	var (
		body = c.Get("body").(*models.CompanyCreatePayload)
	)

	// Process data
	rawData, err := services.CompanyCreate(*body)

	// if err
	if err != nil {
		return ultis.Response400(c, nil, err.Error())
	}

	// Success
	return ultis.Response200(c, rawData, "")
}

// CompanyList to
func CompanyList(c echo.Context) error {
	// Process data
	rawData, err := services.CompanyList()

	// if err
	if err != nil {
		return ultis.Response400(c, nil, err.Error())
	}

	// Success
	return ultis.Response200(c, rawData, "")
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
		return ultis.Response400(c, nil, err.Error())
	}

	// Success
	return ultis.Response200(c, rawData, "")
}

// CompanyUpdate ...
func CompanyUpdate(c echo.Context) error {
	var (
		id           = c.Param("id")
		body         = c.Get("body").(*models.CompanyUpdate)
		companyID, _ = primitive.ObjectIDFromHex(id)
	)

	// Process data
	rawData, err := services.CompanyUpdate(companyID, *body)

	// if err
	if err != nil {
		return ultis.Response400(c, nil, err.Error())
	}

	// Success
	return ultis.Response200(c, rawData, "")
}
