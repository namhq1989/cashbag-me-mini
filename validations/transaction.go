package validations

import (
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"

	"cashbag-me-mini/models"
	"cashbag-me-mini/util"
)

// TransactionCreate ...
func TransactionCreate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			doc models.TransactionCreatePayload
		)

		// ValidateStruct
		c.Bind(&doc)
		_, err := govalidator.ValidateStruct(doc)

		// if err
		if err != nil {
			return util.Response400(c, nil, err.Error())
		}

		// Validate company object id
		companyID, err := util.ValidationObjectID(doc.CompanyID)
		if err != nil {
			return util.Response400(c, nil, err.Error())
		}

		// Validate branch object id
		branchID, err := util.ValidationObjectID(doc.BranchID)
		if err != nil {
			return util.Response400(c, nil, err.Error())
		}

		// Validate user object id
		userID, err := util.ValidationObjectID(doc.UserID)
		if err != nil {
			return util.Response400(c, nil, err.Error())
		}

		// Success
		c.Set("companyID", companyID)
		c.Set("branchID", branchID)
		c.Set("userID", userID)
		c.Set("body", doc)
		return next(c)
	}
}
