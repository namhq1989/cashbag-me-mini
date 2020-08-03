package validations

import (
	"cashbag-me-mini/models"

	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo"
)
//CreateBranch ...
func CreateBranch(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		postBranch := new(models.PostBranch)
		c.Bind(postBranch)
		result, _ := govalidator.ValidateStruct(postBranch)
		if result == true {
			c.Set("body", postBranch)
			next(c)
		}
		return echo.ErrBadRequest
	}
}
