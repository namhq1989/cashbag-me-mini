package routes

import (
	"github.com/labstack/echo/v4"

	"cashbag-me-mini/controllers"
)

// AnalyticChart ...
func AnalyticChart(g *echo.Group) {
	g.GET("/:id", controllers.AnalyticChart)
}