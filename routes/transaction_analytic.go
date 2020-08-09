package routes

import (
	"cashbag-me-mini/controllers"

	"github.com/labstack/echo"
)

// TransactionAnalytic ...
func TransactionAnalytic(g *echo.Group) {
	g.GET("/", controllers.TransactionAnalytic)
}
