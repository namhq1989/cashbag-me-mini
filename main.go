package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"cashbag-me-mini/config"
	"cashbag-me-mini/modules/database"
	"cashbag-me-mini/modules/redis"
	"cashbag-me-mini/modules/zookeeper"
	"cashbag-me-mini/routes"
)

func init() {
	database.Connect("CashBag")
	redis.Connect()
	zookeeper.Connect()
}
func main() {

	cfg := config.GetEnv()
	server := echo.New()

	//CORS
	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTION"},
		MaxAge:           600,
		AllowCredentials: false,
	}))

	// Middleware
	server.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} | ${remote_ip} | ${method} ${uri} - ${status} - ${latency_human}\n",
	}))
	server.Use(middleware.Recover())

	routes.Boostrap(server)
	server.Logger.Fatal(server.Start(cfg.Port))
}
