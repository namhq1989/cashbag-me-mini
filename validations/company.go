package validations

import (
	
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson/primitive"
"log"
	"cashbag-me-mini/models"
	"cashbag-me-mini/dao"
	"cashbag-me-mini/ultis"

)

//CompanyCreate func ...
func CompanyCreate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		doc := new(models.CompanyCreatePayload)
		c.Bind(doc)
		_, err := govalidator.ValidateStruct(doc)

		//if err
		if err != nil{
			return ultis.Response400(c,nil,err.Error())
		}
		
		//Success
		c.Set("body", doc)
		return next(c)
		
	}
}

// CompanyUpdate func ...
func CompanyUpdate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		doc := new(models.CompanyUpdatePayload)
		c.Bind(doc)
		_, err := govalidator.ValidateStruct(doc)

		//if err
		if err != nil{
			return ultis.Response400(c,nil,err.Error())
		}
		//Success
		c.Set("body", doc)
		return next(c)
	}
}

// CompanyCheckID ...
func CompanyCheckID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var(
			id           = c.Param("id")
			companyID, _ = primitive.ObjectIDFromHex(id)
		)
		log.Println(id)

		check := dao.CompanyValidateID(companyID)
		log.Println(check)
		//if err
		if check == false {
			return ultis.Response400(c, nil, "ID khong hop le")
		}

		return next(c)
	}
}