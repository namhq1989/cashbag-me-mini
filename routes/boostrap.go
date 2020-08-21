package routes

import (
	"github.com/labstack/echo/v4"
)

// Boostrap ...
func Boostrap(e *echo.Echo) {
	Branch(e)
	Company(e)
	Transaction(e)
	TransactionAnalytic(e)
	User(e)
	UserProgram(e)
	CompanyAnalytic(e)
	AnalyticChart(e)
}
