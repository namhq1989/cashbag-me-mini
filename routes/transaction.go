package routes

import (
	"cashbag-me-mini/controllers"
	"cashbag-me-mini/validations"
	"github.com/labstack/echo"
)

//Transaction func ...
func Transaction(g *echo.Group) {
	g.POST("", controllers.TransactionCreate, validations.TransactionCreate)
}
