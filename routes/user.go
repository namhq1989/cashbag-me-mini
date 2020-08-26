package routes

import (
	"github.com/labstack/echo/v4"

	"cashbag-me-mini/controllers"
	"cashbag-me-mini/validations"
)

// User ...
func User(e *echo.Echo) {
	routes := e.Group("/users")

	routes.POST("", controllers.UserCreate, validations.UserCreate)
	routes.PUT("/:id", controllers.UserUpdate, validations.UserValidateID, UserCheckExistedByID, validations.UserUpdate)
	routes.GET("/:id/transactions", controllers.TransactionFindByUserID,validations.UserValidateID,UserCheckExistedByID)
}
