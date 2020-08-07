package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"cashbag-me-mini/models"
	"cashbag-me-mini/services"
	"cashbag-me-mini/ultis"
)

// CompanyCreate func to ...
func CompanyCreate(c echo.Context) error {
	var (
		body = c.Get("body").(*models.CompanyCreate)
	)

	//Process data
	rawData, err := services.CompanyCreate(*body)

	//if err
	if err != nil {
		return ultis.Response400(c, nil, err.Error())
	}

	//Success
	return ultis.Response200(c, echo.Map{
		"_id":     rawData.ID,
		"name":    rawData.Name,
		"address": rawData.Address,
	}, "")
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
	return ultis.Response200(c, echo.Map{
		"data": rawData,
	}, "")

}

// CompanyChangeActiveStatus ...
func CompanyChangeActiveStatus(c echo.Context) error {

	id := c.Param("id")

	idCompany, _ := primitive.ObjectIDFromHex(id)
	body := c.Get("body").(*models.CompanyUpdate)
	rawData, err := services.CompanyChangeActiveStatus(idCompany,*body)
	//if err
	if err != nil {
		return ultis.Response400(c, nil, err.Error())
	}

	//success
	return ultis.Response200(c, echo.Map{
		"data": rawData,
	}, "")
}

// PutCompany func ...
func PutCompany(c echo.Context) error {
	id := c.Param("id")
	idCompany, _ := primitive.ObjectIDFromHex(id)
	body := c.Get("body").(*models.CompanyUpdate)
	result := services.PutCompany(idCompany, *body)
	return c.JSON(http.StatusOK, result)

}
