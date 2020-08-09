package routes

import (
	"github.com/labstack/echo/v4"

	"cashbag-me-mini/controllers"
)

// TransactionAnalytic ...
func TransactionAnalytic(g *echo.Group) {
	g.GET("/", controllers.TransactionAnalyticList)
}
