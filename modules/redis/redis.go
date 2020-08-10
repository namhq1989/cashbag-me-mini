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

	// Ping Redis
	pong, err := rdb.Ping().Result()
	fmt.Println(pong, err)
	db = rdb
}

// GetUser ...
func GetUser() string {
	value := db.Get("user").Val()
	return value
}

// SetUser ...
func SetUser(value string) {
	const setTime = 30000000000

	err := db.Set("user", value, setTime).Err()
	if err != nil {
		log.Println(err)
	}
}
