package routes

import (
	"cashbag-me-mini/controllers"

	"github.com/labstack/echo"
)

// TranAnalytic ...
func TranAnalytic(g *echo.Group) {
	//Method Get
	g.GET("/", controllers.TranAnalytic)
}
