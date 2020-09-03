package validations

import (
	"github.com/labstack/echo/v4"

	"cashbag-me-mini/util"
)

// UserValidateID ...
func UserValidateID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			id          = c.Param("id")
			userID, err = util.ValidationObjectID(id)
		)

		// if err
		if err != nil {
			return util.Response400(c, nil, "ID khong hop le")
		}

		c.Set("userID", userID)
		return next(c)
	}
}
