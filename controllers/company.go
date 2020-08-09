package controllers

import (

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"cashbag-me-mini/models"
	"cashbag-me-mini/services"
	"cashbag-me-mini/ultis"
)

// CompanyCreate func to ...
func CompanyCreate(c echo.Context) error {
	var (
		body = c.Get("body").(*models.CompanyCreatePayload)
	)

	//Process data
	rawData, err := services.CompanyCreate(*body)

	//if err
	if err != nil {
		return ultis.Response400(c, nil, err.Error())
	}

	//Success
	return ultis.Response200(c, rawData, "")
}

// CompanyList to
func CompanyList(c echo.Context) error {
	//Process data
	rawData, err := services.CompanyList()

	//if err
	if err != nil {
		return ultis.Response400(c, nil, err.Error())
	}

	//success
	return ultis.Response200(c, rawData, "")

}

// CompanyChangeActiveStatus ...
func CompanyChangeActiveStatus(c echo.Context) error {
	var (
		id           = c.Param("id")
		companyID, _ = primitive.ObjectIDFromHex(id)
	)


	rawData, err := services.CompanyChangeActiveStatus(companyID)

	//if err
	if err != nil {
		return ultis.Response400(c, nil, err.Error())
	}

	//success
	return ultis.Response200(c,rawData, "")
}

// CompanyUpdate func ...
func CompanyUpdate(c echo.Context) error {
	var (
		id           = c.Param("id")
		companyID, _ = primitive.ObjectIDFromHex(id)
		body = c.Get("body").(*models.CompanyUpdatePayload)

	)

	rawData, err := services.CompanyUpdate(companyID, *body)

	if err != nil {
		return ultis.Response400(c, nil, err.Error())
	}

	//success
	return ultis.Response200(c, rawData, "")
}
