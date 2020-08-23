package validations

import (
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"

	"cashbag-me-mini/dao"
	"cashbag-me-mini/models"
	"cashbag-me-mini/util"
)

// UserProgramCreate ...
func UserProgramCreate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			doc models.UserProgramCreatePayload
		)

		// ValidateStruct
		c.Bind(&doc)
		_, err := govalidator.ValidateStruct(doc)

		//if err
		if err != nil {
			return util.Response400(c, nil, err.Error())
		}

		// Validate object id
		companyID, err := util.ValidationObjectID(doc.CompanyID)
		if err != nil {
			return util.Response400(c, nil, err.Error())
		}

		// Validate existed in db
		company, err := dao.CompanyFindByID(companyID)
		if company.ID.IsZero() {
			return util.Response400(c, nil, "Khong tim thay Cong Ty")
		}

		if err != nil {
			return util.Response400(c, nil, err.Error())
		}

		//Success
		c.Set("body", doc)
		return next(c)
	}
}
