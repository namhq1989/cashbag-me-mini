package zookeeper

import (
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

var ZDB *zk.Conn

//ConnectZookeeper ...
func ConnectZookeeper() {
	c, _, err := zk.Connect([]string{"zookeeper:2181"}, time.Second)
	if err != nil {
		panic(err)
	}
	ZDB = c
}

// //GetValueFromZoo ...
// func GetValueFromZoo(path string) string{
// 	res,_,_:=Zoo.Get(path)
// 	return string(res)
//}
