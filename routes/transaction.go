package routes

import (
	"github.com/labstack/echo/v4"

	"cashbag-me-mini/controllers"
	"cashbag-me-mini/validations"
)

//Transaction func ...
func Transaction(g *echo.Group) {
	g.POST("", controllers.TransactionCreate, validations.TransactionCreate)
}
