package routes

import(
	"github.com/labstack/echo"
)
// Boostrap ...
func Boostrap(e *echo.Echo)  {
	Branch(e.Group("/branches"))
	Company(e.Group("/companies"))
	Transaction(e.Group("/transactions"))
	TranAnalytic(e.Group("/transaction_analytic"))
}