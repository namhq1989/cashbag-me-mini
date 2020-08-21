package routes

import (
	"github.com/labstack/echo/v4"

	"cashbag-me-mini/controllers"
	"cashbag-me-mini/validations"
)

//Transaction func ...
func Transaction(e *echo.Echo) {
	routes := e.Group("/transactions")

	routes.POST("", controllers.TransactionCreate, validations.TransactionCreate)
}
