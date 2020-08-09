package redis

import (
	"fmt"
	"log"

	"github.com/go-redis/redis"
	"cashbag-me-mini/config"
)

var (
	RDB *redis.Client
)

//Connect ...
func Connect() {
	envVars := config.GetEnv()
	rdb := redis.NewClient(&redis.Options{
		Addr:    envVars.RedisURI, // use default Addr
		Password: envVars.RedisPass,               // no password set
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

const setTime =30000000000
//SetValueRedis ...
func SetValueRedis(key string, value string) {
	err := RDB.Set(key, value, setTime).Err()
	if err != nil {
		log.Println(err)
	}
}
