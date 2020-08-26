package routes

import (
	"github.com/labstack/echo/v4"

	"cashbag-me-mini/controllers"
	"cashbag-me-mini/validations"
)

// Branch ...
func Branch(e *echo.Echo) {
	routes := e.Group("/branches")

	routes.GET("", controllers.BranchList)
	routes.POST("", controllers.BranchCreate, validations.BranchCreate, CompanyCheckExistedByID)
	routes.PATCH("/:id/active", controllers.BranchChangeActiveStatus, validations.BranchValidateID, BranchCheckExistedByID)
	routes.PUT("/:id", controllers.BranchUpdate, validations.BranchValidateID, BranchCheckExistedByID, validations.BranchUpdate)
}
