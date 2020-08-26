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
	routes.POST("", controllers.BranchCreate, validations.BranchCreate, companyCheckExistedByID)
	routes.PATCH("/:id/active", controllers.BranchChangeActiveStatus, validations.BranchValidateID, branchCheckExistedByID)
	routes.PUT("/:id", controllers.BranchUpdate, validations.BranchValidateID, branchCheckExistedByID, validations.BranchUpdate)
}
