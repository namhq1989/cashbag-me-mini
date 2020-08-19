package routes

import (
	"github.com/labstack/echo/v4"

	"cashbag-me-mini/controllers"
)

// CompanyAnalytic ...
func CompanyAnalytic(g *echo.Group) {
	g.GET("", controllers.CompanyAnalyticList)
}