package validations

import (
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	
	"cashbag-me-mini/dao"
	"cashbag-me-mini/models"
	"cashbag-me-mini/ultis" 
	
)

// UserCreate ...
func UserCreate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			doc = new(models.UserCreatePayload)
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


