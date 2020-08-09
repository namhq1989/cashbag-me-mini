package routes

import (
	"cashbag-me-mini/controllers"
	"cashbag-me-mini/validations"

	"github.com/labstack/echo"
)

// Branch ...
func Branch(g *echo.Group) {
	g.GET("", controllers.BranchList)
	g.POST("", controllers.BranchCreate, validations.BranchCreate)
	g.PATCH("/:id", controllers.BranchChangeActiveStatus, validations.BranchCheckID)
	g.PUT("/:id", controllers.PutBranch,validations.BranchCheckID, validations.BranchUpdate)
}
