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
 
//UpdateBranch func ..
func UpdateBranch(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		putBranch := new(models.PutBranch)
		c.Bind(putBranch)
		result, _ := govalidator.ValidateStruct(putBranch)
		if result == true {
			c.Set("body", putBranch)
			next(c)
		}
		return echo.ErrBadRequest
	}
}


