package main
<<<<<<< HEAD

import (
	"cashbag-me-mini/modules/database"
	"cashbag-me-mini/routes"

	"github.com/labstack/echo"
)

=======
import (
	"cashbag-me-mini/modules/database"
	"cashbag-me-mini/routes"
	"github.com/labstack/echo"
)
>>>>>>> c4e7b26aca10bbf49db1183a063730b16411c8e6
func init() {
	database.Connectdb("CashBag")
}
func main() {
	server := echo.New()
	routes.CompanyRoute(server.Group("/companies"))
<<<<<<< HEAD
	routes.Branch(server.Group("/branchs"))
=======
>>>>>>> c4e7b26aca10bbf49db1183a063730b16411c8e6
	server.Logger.Fatal(server.Start(":8080"))
}
