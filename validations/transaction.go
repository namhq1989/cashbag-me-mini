package validations

import (
	"cashbag-me-mini/dao"

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
		userID,err :=util.ValidationObjectID(doc.UserID)
		if err != nil {
			return util.Response400(c, nil, err.Error())
		}

		// Validate company existed in db
		company, err := dao.CompanyFindByID(companyID)
		if company.ID.IsZero() {
			return util.Response400(c, nil, "Khong tim thay Cong Ty ")
		}
		if err != nil {
			return util.Response400(c, nil, err.Error())
		}

		// Validate branch existed in db
		branch, err := dao.BranchFindByID(branchID)
		if branch.ID.IsZero() {
			return util.Response400(c, nil, "Khong tim thay Chi Nhanh")
		}
		if err != nil {
			return util.Response400(c, nil, err.Error())
		}

		// Validate user existed in db
		user, err := dao.UserFindByID(userID)
		if user.ID.IsZero() {
			return util.Response400(c, nil, "Khong tim thay Nguoi Dung")
		}
		if err != nil {
			return util.Response400(c, nil, err.Error())
		}


		// Success
		c.Set("body", doc)
		return next(c)
	}
}
