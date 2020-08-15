package routes

import(
	"github.com/labstack/echo/v4"

	"cashbag-me-mini/controllers"
	"cashbag-me-mini/validations"
)

// UserProgram ....
func UserProgram(g *echo.Group) {
	g.POST("",controllers.UserProgramCreate,validations.UserProgramCreate)
}