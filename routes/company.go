package routes

import (
	"cashbag-me-mini/controllers"
	"cashbag-me-mini/validations"

	"github.com/labstack/echo"
)

//Company  to ...
func Company(g *echo.Group) {

	g.POST("", controllers.CompanyCreate, validations.CompanyCreate)
	g.GET("", controllers.CompanyList)
	g.PATCH("/:id", controllers.CompanyChangeActiveStatus)
	g.PUT("/:id", controllers.PutCompany, validations.CompanyUpdate)

}
