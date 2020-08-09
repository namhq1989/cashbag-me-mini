package routes

import (
	"github.com/labstack/echo/v4"

	"cashbag-me-mini/controllers"
	"cashbag-me-mini/validations"
)

// Branch ...
func Branch(g *echo.Group) {
	g.GET("", controllers.BranchList)
	g.POST("", controllers.BranchCreate, validations.BranchCreate)
	g.PATCH("/:id", controllers.BranchChangeActiveStatus, validations.BranchValidateID)
	g.PUT("/:id", controllers.BranchUpdate, validations.BranchValidateID, validations.BranchUpdate)
}
