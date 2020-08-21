package routes

import (
	"github.com/labstack/echo/v4"

	"cashbag-me-mini/controllers"
)

// CompanyAnalytic ...
func CompanyAnalytic(e *echo.Echo) {
	routes := e.Group("/company-analytics")

	routes.GET("", controllers.CompanyAnalyticList)
}
