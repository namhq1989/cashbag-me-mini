package zookeeper

import (
	"cashbag-me-mini/config"
	"os"

	"github.com/samuel/go-zookeeper/zk"

	"fmt"
	"time"
)

var conn *zk.Conn

// Connect ...
func Connect() {
	var (
		uri     = os.Getenv("ZOOKEEPER_URI")
		envVars = config.GetEnv()
	)
	conn, _, err := zk.Connect([]string{uri}, time.Second*30)
	if err != nil {
		fmt.Println("ZookeeperURI:", uri)
		panic(err)
	}
	fmt.Println("Zookeeper Connected to", uri)

	// Get env key
	// App port
	appPort, _, _ := conn.Get("/app/port")
	envVars.AppPort = string(appPort)

	// Database
	databaseURI, _, _ := conn.Get("/database/uri")
	envVars.Database.URI = string(databaseURI)
	databaseName, _, _ := conn.Get("/database/name")
	envVars.Database.Name = string(databaseName)
	databaseTestName, _, _ := conn.Get("/database/test/name")
	envVars.Database.TestName = string(databaseTestName)

	// Redis
	redisURI, _, _ := conn.Get("/redis/uri")
	envVars.Redis.URI = string(redisURI)
	redisPass, _, _ := conn.Get("/redis/pass")
	envVars.Redis.Pass = string(redisPass)

	// Users
	userString, _, _ := conn.Get("/users")
	envVars.UserString = string(userString)
}
