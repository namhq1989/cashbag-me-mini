package routes

import (
	"cashbag-me-mini/controllers"
	"cashbag-me-mini/validations"

	"github.com/labstack/echo"
)

// Branch ...
func Branch(g *echo.Group) {
	//Method Get
	g.GET("", controllers.ListBranch)
	//Method Post
	g.POST("", controllers.CreateBranch, validations.CreateBranch)

}
