package routes

import (
	"cashbag-me-mini/controllers"
	"cashbag-me-mini/validations"
	"github.com/labstack/echo"
)

//TransactionRoute func ...
func TransactionRoute(g *echo.Group) {
	g.POST("", controllers.CreateTransaction, validations.CreateTransaction)
}
