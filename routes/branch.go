package routes

import (
	"cashbag-me-mini/controllers"

	"github.com/labstack/echo"
)

// Branch ...
func Branch(g *echo.Group) {
	//Method Get
	g.GET("", controllers.ListBranch)
}
