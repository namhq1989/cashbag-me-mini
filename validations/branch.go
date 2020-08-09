package validations

import (
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"

	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"cashbag-me-mini/dao"
	"cashbag-me-mini/models"
	"cashbag-me-mini/ultis"
)

// BranchCreate ...
func BranchCreate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			doc = new(models.BranchCreatePayload)
		)

		// ValidateStruct
		c.Bind(doc)
		log.Println(doc)

		_, err := govalidator.ValidateStruct(doc)
		log.Println(err)

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
		var (
			doc = new(models.BranchUpdateBPayload)
		)

		// ValidateStruct
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
			return ultis.Response400(c, nil, "ID khong hop le")
		}

		return next(c)
	}
}
