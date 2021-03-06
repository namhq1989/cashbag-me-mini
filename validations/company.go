package validations

import (
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"

	"cashbag-me-mini/models"
	"cashbag-me-mini/util"
)

// CompanyCreate ...
func CompanyCreate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			doc models.CompanyCreatePayload
		)

		// ValidateStruct
		c.Bind(&doc)
		_, err := govalidator.ValidateStruct(doc)

		//if err
		if err != nil {
			return util.Response400(c, nil, err.Error())
		}

		//Success
		c.Set("body", doc)

		return next(c)
	}
}

// CompanyUpdate ...
func CompanyUpdate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			doc models.CompanyUpdatePayload
		)

		// ValidateStruct
		c.Bind(&doc)
		_, err := govalidator.ValidateStruct(doc)

		//if err
		if err != nil {
			return util.Response400(c, nil, err.Error())
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
			id             = c.Param("id")
			companyID, err = util.ValidationObjectID(id)
		)

		// if err
		if err != nil {
			return util.Response400(c, nil, "ID company khong hop le ")

		}

		c.Set("companyID", companyID)

		return next(c)
	}
}
