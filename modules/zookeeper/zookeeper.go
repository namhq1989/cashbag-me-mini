package zookeeper

import (
	"fmt"
	"os"
	"time"

	"github.com/samuel/go-zookeeper/zk"

	"cashbag-me-mini/config"
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
	appTransactionPort, _, _ := conn.Get("/app/port/transaction")
	envVars.AppTransactionPort = string(appTransactionPort)

	// Database
	databaseURI, _, _ := conn.Get("/database/uri")
	envVars.Database.URI = string(databaseURI)
	databaseTransactionName, _, _ := conn.Get("/database/name/transaction")
	envVars.Database.TransactionName = string(databaseTransactionName)
	databaseTestName, _, _ := conn.Get("/database/test/name")
	envVars.Database.TestName = string(databaseTestName)

	// Redis
	redisURI, _, _ := conn.Get("/redis/uri")
	envVars.Redis.URI = string(redisURI)
	redisPass, _, _ := conn.Get("/redis/pass")
	envVars.Redis.Pass = string(redisPass)

	// GRPC
	grpcURI, _, _ := conn.Get("/grpc/uri")
	envVars.GRPCUri = string(grpcURI)
}
