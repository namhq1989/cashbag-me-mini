package routes

import (
	"cashbag-me-mini/controllers"
	"cashbag-me-mini/validations"

	"github.com/labstack/echo"
)

//CompanyRoute to ...
func CompanyRoute(g *echo.Group) {

	
	g.POST("", controllers.CreateCompany, validations.CreateCompany)
	
}
