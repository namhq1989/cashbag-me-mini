package redis

import (
	"github.com/go-redis/redis"

	"fmt"
	"log"

	"cashbag-me-mini/config"
)

var (
	db *redis.Client
)

//Connect ...
func Connect() {
	envVars := config.GetEnv()
	rdb := redis.NewClient(&redis.Options{
		Addr:     envVars.RedisURI,  // use default Addr
		Password: envVars.RedisPass, // no password set
		DB:       0,                 // use default DB
	})
	pong, err := rdb.Ping().Result()
	fmt.Println(pong, err)
	db = rdb
}

//GetValueRedis ...
func GetValueRedis(key string) string {
	value := db.Get(key).Val()
	return value
}

const setTime = 30000000000

//SetValueRedis ...
func SetValueRedis(key string, value string) {
	err := db.Set(key, value, setTime).Err()
	if err != nil {
		log.Println(err)
	}
}
