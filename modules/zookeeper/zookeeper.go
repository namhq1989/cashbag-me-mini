package zookeeper

import (
	"fmt"
	"time"

	"github.com/samuel/go-zookeeper/zk"
	"cashbag-me-mini/config"
)

var ZDB *zk.Conn

//Connect ...
func Connect() {
	envVars := config.GetEnv()
	c, _, err := zk.Connect([]string{envVars.ZookeeperURI}, time.Second*30)
	if err != nil {
		panic(err)
	}
	ZDB = c
	fmt.Println("Connect Zookeeper")
}

//GetValueFromZoo ...
func GetValueFromZoo(path string) string {
	res, _, _ := ZDB.Get(path)
	return string(res)
}
