package redis

import (
	"fmt"
	"log"

	"github.com/go-redis/redis"
)

var (
	RDB *redis.Client
)

//ConnectRDB ...
func ConnectRDB() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // use default Addr
		Password: "",               // no password set
		DB:       0,                // use default DB
	})
	pong, err := rdb.Ping().Result()
	fmt.Println(pong, err)
	RDB=rdb
}

//GetValueRedis ...
func GetValueRedis(key string) string {
	value:= RDB.Get(key).Val()
	return value
}

//SetValueRedis ...
func SetValueRedis(key string, value string) {
	err := RDB.Set(key, value, 30000000000).Err()
	if err != nil {
		log.Println(err)
	}
}
