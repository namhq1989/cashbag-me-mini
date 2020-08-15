package validations

import (
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"cashbag-me-mini/dao"
	"cashbag-me-mini/models"
	"cashbag-me-mini/util"
)

// BranchCreate ...
func BranchCreate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			doc = new(models.BranchCreatePayload)
		)

		// ValidateStruct
		c.Bind(doc)

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

		// Validate existed in db
		company, err := dao.CompanyFindByID(companyID)
		if err != nil {
			return util.Response400(c, nil, err.Error())
		}
		if company.ID.IsZero() {
			return util.Response400(c, nil, "Khong tim thay Cong Ty")
		}

		// Success
		c.Set("body", doc)
		return next(c)
	}
}

// BranchUpdate ...
func BranchUpdate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			doc = new(models.BranchUpdateBPayload)
		)

		// ValidateStruct
		c.Bind(doc)
		_, err := govalidator.ValidateStruct(doc)

		//if err
		if err != nil {
			return util.Response400(c, nil, err.Error())
		}

		// Validate object id
		id := c.Param("id")
		branchID, err := util.ValidationObjectID(id)
		if err != nil {
			return util.Response400(c, nil, err.Error())
		}

		// Validate existed in db
		branch, err := dao.BranchFindByID(branchID)
		if err != nil {
			return util.Response400(c, nil, err.Error())
		}
		if branch.ID.IsZero() {
			return util.Response400(c, nil, "Khong tim thay Branch")
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
			id          = c.Param("id")
			branchID, _ = primitive.ObjectIDFromHex(id)
			branch, _   = dao.BranchFindByID(branchID)
		)

		// Validate ID
		if branch.ID.IsZero() {
			return util.Response400(c, nil, "ID khong hop le")
		}

		return next(c)
	}
}
