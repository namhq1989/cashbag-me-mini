package validations

import (
	"cashbag-me-mini/models"
	"cashbag-me-mini/ultis"

	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
)

// CompanyCreate ...
func CompanyCreate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		doc := new(models.PayloadOfCreateCompany)
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
		doc := new(models.PayloadOfUpdateCompany)
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
