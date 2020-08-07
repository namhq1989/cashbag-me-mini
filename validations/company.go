package validations

import (
	

	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo"

	"cashbag-me-mini/models"
	"cashbag-me-mini/ultis"
)

//CompanyCreate func ...
func CompanyCreate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		doc := new(models.CompanyCreate)
		c.Bind(doc)
		_, err := govalidator.ValidateStruct(doc)

		//if err
		if err != nil{
			return ultis.Response400(c,nil,err.Error())
		}
		
		//Success
		c.Set("body", doc)
		return next(c)
		
	}
}

// CompanyUpdate func ...
func CompanyUpdate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		doc := new(models.CompanyUpdate)
		c.Bind(doc)
		_, err := govalidator.ValidateStruct(doc)

		//if err
		if err != nil{
			return ultis.Response400(c,nil,err.Error())
		}
		//Success
		c.Set("body", doc)
		return next(c)
	}
}
