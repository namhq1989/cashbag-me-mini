package validations

import (
	"cashbag-me-mini/dao"
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"

	"cashbag-me-mini/models"
	"cashbag-me-mini/ultis"
)

// BranchCreate ...
func BranchCreate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		doc := new(models.BranchCreatePayload)
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

// BranchUpdate ...
func BranchUpdate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		doc := new(models.BranchUpdateBPayload)
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

// BranchCheckID ...
func BranchCheckID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var(
			branchID          = c.Param("id")
		)

		check := dao.BranchValidateID(branchID)

		//if err
		if check == false {
			return ultis.Response400(c, nil, "ID khong hop le")
		}

		return next(c)
	}
}

