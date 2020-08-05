package zookeeper

import (
	"fmt"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

var ZDB *zk.Conn

//ConnectZookeeper ...
func ConnectZookeeper() {
	c, _, err := zk.Connect([]string{"127.0.0.1:2181"}, time.Second)
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
