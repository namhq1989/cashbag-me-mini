package validations

import (
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"

	"cashbag-me-mini/models"
	"cashbag-me-mini/util"
)

// BranchCreate ...
func BranchCreate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			doc models.BranchCreatePayload
		)

		// ValidateStruct
		c.Bind(&doc)

		_, err := govalidator.ValidateStruct(doc)

		// if err
		if err != nil {
			return util.Response400(c, nil, err.Error())
		}

		// Validate object id
		companyID, err := util.ValidationObjectID(doc.CompanyID)
		if err != nil {
			return util.Response400(c, nil, err.Error())
		}

		// Success
		c.Set("body", doc)
		c.Set("companyID", companyID)
		return next(c)
	}
}

// BranchUpdate ...
func BranchUpdate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			doc models.BranchUpdatePayload
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

// BranchValidateID ...
func BranchValidateID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			id            = c.Param("id")
			branchID, err = util.ValidationObjectID(id)
		)

		// if err
		if err != nil {
			return util.Response400(c, nil, "ID branch khong hop le ")
		}

		c.Set("branchID", branchID)

		return next(c)
	}
}
