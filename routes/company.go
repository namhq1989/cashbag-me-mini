package routes

import (
	
	"github.com/labstack/echo"

	"cashbag-me-mini/controllers"
	"cashbag-me-mini/validations"
)


//Company  to ...
func Company(g *echo.Group) {

	g.POST("", controllers.CompanyCreate, validations.CompanyCreate)
	g.GET("", controllers.CompanyList)
	g.PATCH("/:id", controllers.CompanyChangeActiveStatus)
	g.PUT("/:id", controllers.CompanyUpdate, validations.CompanyUpdate)

}
