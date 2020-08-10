package routes

import (
	"github.com/labstack/echo/v4"

	"cashbag-me-mini/controllers"
	"cashbag-me-mini/validations"
)

// Company ...
func Company(g *echo.Group) {
	g.POST("", controllers.CompanyCreate, validations.CompanyCreate)
	g.GET("", controllers.CompanyList)
	g.PATCH("/:id/active", controllers.CompanyChangeActiveStatus, validations.CompanyValidateID)
	g.PUT("/:id", controllers.CompanyUpdate, validations.CompanyValidateID, validations.CompanyUpdate)
}
