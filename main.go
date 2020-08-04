package main

import (
	"cashbag-me-mini/modules/database"
	"cashbag-me-mini/routes"

	"github.com/labstack/echo"
)

func init() {
	database.Connectdb("CashBag")
}
func main() {
	server := echo.New()
	routes.Branch(server.Group("/branches"))
	routes.CompanyRoute(server.Group("/companies"))
	server.Logger.Fatal(server.Start(":8080"))
}
