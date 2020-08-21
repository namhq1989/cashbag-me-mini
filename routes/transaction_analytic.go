package routes

import (
	"github.com/labstack/echo/v4"

	"cashbag-me-mini/controllers"
)

// TransactionAnalytic ...
func TransactionAnalytic(e *echo.Echo) {
	routes := e.Group("/transaction-analytics")

	routes.GET("/", controllers.TransactionAnalyticList)
}
