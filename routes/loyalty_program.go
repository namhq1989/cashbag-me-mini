package routes

import (
	"github.com/labstack/echo/v4"

	"cashbag-me-mini/controllers"
	"cashbag-me-mini/validations"
)

//UserProgram func ...
func UserProgram(e *echo.Echo) {
	routes := e.Group("/loyalty-programs")

	routes.POST("", controllers.LoyaltyProgramCreate, validations.LoyaltyProgramCreate, CompanyCheckExistedByID)
}
