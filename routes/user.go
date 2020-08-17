package routes

import (
	"github.com/labstack/echo/v4"

	"cashbag-me-mini/controllers"
	"cashbag-me-mini/validations"
)

// User ...
func User(g *echo.Group) {
	g.POST("", controllers.UserCreate,validations.UserCreate)
	g.PUT("/:id",controllers.UserUpdate, validations.UserValidateID, validations.UserUpdate)
}	