package zookeeper

import (
	"github.com/samuel/go-zookeeper/zk"

	"fmt"
	"time"

	"cashbag-me-mini/config"
)

var db *zk.Conn

//Connect ...
func Connect() {
	envVars := config.GetEnv()
	c, _, err := zk.Connect([]string{envVars.ZookeeperURI}, time.Second*30)
	if err != nil {
		panic(err)
	}
	db = c
	fmt.Println("Connect Zookeeper")
}

//GetUser ...
func GetUser() string {
	res, _, _ := db.Get("/users")
	return string(res)
}
