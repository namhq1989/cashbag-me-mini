package main

import (
	"cashbag-me-mini/modules/database"
	"cashbag-me-mini/modules/redis"
	"cashbag-me-mini/routes"

	"github.com/labstack/echo"
)

func init() {
	database.Connectdb("CashBag")
	redis.ConnectRDB()
}
func main() {
	server := echo.New()
	routes.Branch(server.Group("/branches"))
	routes.CompanyRoute(server.Group("/companies"))
	routes.TranAnalytic(server.Group("/tranAnalytic"))
	server.Logger.Fatal(server.Start(":8080"))
}
