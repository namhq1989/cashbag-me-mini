package routes

import (
	"github.com/labstack/echo/v4"
)

// Boostrap ...
func Boostrap(e *echo.Echo) {
	Branch(e.Group("/branches"))
	Company(e.Group("/companies"))
	Transaction(e.Group("/transactions"))
	TransactionAnalytic(e.Group("/transaction-analytics"))
	User(e.Group("/user"))
	UserProgram(e.Group("/user-program"))
}
