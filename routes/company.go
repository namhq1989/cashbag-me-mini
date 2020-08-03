package routes

import (
	"cashbag-me-mini/controllers"

	"github.com/labstack/echo"
)

//CompanyRoute to ...
func CompanyRoute(g *echo.Group) {

	g.GET("", controllers.CompanyList)
	g.POST("", controllers.CreateCompany)
	g.PUT("/:id", controllers.CompanyUpdate)
	g.PATCH("/:id", controllers.CompanyActive)
}
