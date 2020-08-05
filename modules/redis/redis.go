package redis

import (
	"fmt"

	"github.com/go-redis/redis"
)

var (
	RDB *redis.Client
)

//ConnectRDB ...
func ConnectRDB(){
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // use default Addr
		Password: "",               // no password set
		DB:       0,                // use default DB
	})	
	pong, err := rdb.Ping().Result()
	fmt.Println(pong, err)
}
