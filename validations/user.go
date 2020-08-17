package validations

import (
	

	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"cashbag-me-mini/dao"
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

		// Validate object id
		id := c.Param("id")
		userID, _ := util.ValidationObjectID(id)
		if err != nil {
			return util.Response400(c, nil, err.Error())
		}

		// Validate existed in db
		user, _ := dao.UserFindByID(userID)
		if user.ID.IsZero() {
			return util.Response400(c, nil, "Khong tim thay User")
		}
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
			id        = c.Param("id")
			userID, _ = primitive.ObjectIDFromHex(id)
			user, _   = dao.UserFindByID(userID)
		)

		// Validate ID
		if user.ID.IsZero() {
			return util.Response400(c, nil, "ID khong hop le")
		}

		return next(c)
	}
}
