package validations

import (
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"cashbag-me-mini/dao"
	"cashbag-me-mini/models"
	"cashbag-me-mini/dao"
	"cashbag-me-mini/ultis"

)

// CompanyCreate ...
func CompanyCreate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			doc = new(models.CompanyCreatePayload)
		)

		// ValidateStruct
		c.Bind(doc)
		_, err := govalidator.ValidateStruct(doc)

		//if err
		if err != nil {
			return ultis.Response400(c, nil, err.Error())
		}

		//Success
		c.Set("body", doc)
		return next(c)
	}
}

// CompanyUpdate func ...
func CompanyUpdate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			doc = new(models.CompanyUpdatePayload)
		)

		// ValidateStruct
		c.Bind(doc)
		_, err := govalidator.ValidateStruct(doc)

		//if err
		if err != nil {
			return ultis.Response400(c, nil, err.Error())
		}

		//Success
		c.Set("body", doc)
		return next(c)
	}
}

// CompanyValidateID ...
func CompanyValidateID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			id           = c.Param("id")
			companyID, _ = primitive.ObjectIDFromHex(id)
			company, _   = dao.CompanyFindByID(companyID)
		)

		// Validate ID
		if company.ID.IsZero() {
			return ultis.Response400(c, nil, "ID khong hop le")
		}

		return next(c)
	}
}
