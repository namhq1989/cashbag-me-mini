package routes

import (
	"cashbag-me-mini/controllers"

	"github.com/labstack/echo"
)

//CompanyRoute to ...
func CompanyRoute(g *echo.Group) {

	g.GET("", controllers.CompanyList)

}
