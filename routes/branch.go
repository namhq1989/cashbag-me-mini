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
	routes.POST("", controllers.BranchCreate, validations.BranchCreate)
	routes.PATCH("/:id/active", controllers.BranchChangeActiveStatus, validations.BranchValidateID,Checkext,CheckExit)
	routes.PUT("/:id", controllers.BranchUpdate, validations.BranchValidateID, validations.BranchUpdate)
}
