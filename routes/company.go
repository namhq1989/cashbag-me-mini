package routes

import (
	"github.com/labstack/echo/v4"

	"cashbag-me-mini/controllers"
	"cashbag-me-mini/validations"
)

// Company ...
func Company(e *echo.Echo) {
	routes := e.Group("/companies")

	routes.POST("", controllers.CompanyCreate, validations.CompanyCreate)
	routes.GET("", controllers.CompanyList)
	routes.PATCH("/:id/active", controllers.CompanyChangeActiveStatus,validations.CompanyValidateID,companyCheckExistedByID)
	routes.PUT("/:id", controllers.CompanyUpdate, validations.CompanyValidateID,companyCheckExistedByID, validations.CompanyUpdate)
}

