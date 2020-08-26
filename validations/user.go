package validations

import (
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"

	"cashbag-me-mini/models"
	"cashbag-me-mini/util"
)

// UserCreate ...
func UserCreate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			doc models.UserCreatePayload
		)

		// ValidateStruct
		c.Bind(&doc)
		_, err := govalidator.ValidateStruct(doc)

		// if err
		if err != nil {
			return util.Response400(c, nil, err.Error())
		}

		// Success
		c.Set("body", doc)
		return next(c)
	}
}

// UserUpdate ...
func UserUpdate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			doc models.UserUpdatePayload
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
