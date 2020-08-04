package validations

import (
	"cashbag-me-mini/models"

	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo"
)

//CreateCompany func ...
func CreateCompany(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		postCompany := new(models.PostCompany)
		c.Bind(postCompany)
		result, _ := govalidator.ValidateStruct(postCompany)
		if result == true {
			c.Set("body", postCompany)
			next(c)
		}
		return echo.ErrBadRequest
	}
}

//PutCompany func ...
func PutCompany(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		putCompany := new(models.PutCompany)
		c.Bind(putCompany)
		result, _ := govalidator.ValidateStruct(putCompany)
		if result == true {
			c.Set("body", putCompany)
			next(c)
		}
		return echo.ErrBadRequest
	}
}
