package validations

import (
	"cashbag-me-mini/models"

	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo"
)

//CreateTransaction  ...
func CreateTransaction(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		postTransaction := new(models.PostTransaction)
		c.Bind(postTransaction)
		result, _ := govalidator.ValidateStruct(postTransaction)
		if result == true {
			c.Set("body", postTransaction)
			next(c)
		}
		return echo.ErrBadRequest
	}
}
