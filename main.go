package main

import (
	"github.com/labstack/echo"
	"cashbag-me-mini/routes"
	"cashbag-me-mini/modules/database"
)
func init() {
	database.Connectdb("CashBag")
}
func main() {
	server := echo.New()
	routes.Branch(server.Group("/branchs"))
	server.Logger.Fatal(server.Start(":8080"))
}
