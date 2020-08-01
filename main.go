package main

import (
	"github.com/labstack/echo"
	"cashbag-me-mini/routes"
)
func main() {
	server := echo.New()
	//routes.Company(server.Group("/companys"))
	routes.Branch(server.Group("/branchs"))
	server.Logger.Fatal(server.Start(":8080"))
}
