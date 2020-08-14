package routes

import (
	"github.com/labstack/echo/v4"

	"cashbag-me-mini/controllers"
	
)

// User ...
func User(g *echo.Group) {
	g.POST("", controllers.UserCreate)
	g.PUT("/:id",controllers.UserUpdate)
}