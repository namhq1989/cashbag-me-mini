package validations

import (
	"cashbag-me-mini/models"

	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo"
)

// TransactionCreate ...
func TransactionCreate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		doc := new(models.TransactionCreatePayload)
		c.Bind(doc)
		result, _ := govalidator.ValidateStruct(doc)
		if result == true {
			c.Set("body", doc)
			next(c)
		}
		return echo.ErrBadRequest
	}
}
