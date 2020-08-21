package routes

import (
	"github.com/labstack/echo/v4"

	"cashbag-me-mini/controllers"
)

// AnalyticChart ...
func AnalyticChart(e *echo.Echo) {
	routes := e.Group("/analytic-charts")
	
	routes.GET("/:id", controllers.AnalyticChart)
}
