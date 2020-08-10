package config

import (
	"os"
)

// ENV ...
type ENV struct {
	IsDev            bool
	ZookeeperURI     string
	ZookeeperTestURI string
	DatabaseURI      string
	DatabaseName     string
	DatabaseTestName string
	RedisURI         string
	RedisPass        string
	Port             string
}

var env ENV

func initENV() {
	os.Setenv("DATABASE_URI", "mongodb+srv://CashbagMe:Cashbag@cluster0.epe8y.gcp.mongodb.net/Cluster0?retryWrites=true&w=majority")
	os.Setenv("DATABASE_NAME", "CashBag")
	os.Setenv("DATABASE_TEST_NAME", "CashBag-test")
	os.Setenv("ZOOKEEPER_URI", "zookeeper:2181")
	os.Setenv("ZOOKEEPER_TEST_URI", "127.0.0.1:2181")
	os.Setenv("REDIS_URI", "redis:6379")
	os.Setenv("REDIS_PASS", "")
	os.Setenv("PORT", ":8080")

	env = ENV{
		IsDev:            true,
		DatabaseURI:      os.Getenv("DATABASE_URI"),
		DatabaseName:     os.Getenv("DATABASE_NAME"),
		DatabaseTestName: os.Getenv("DATABASE_TEST_NAME"),
		ZookeeperURI:     os.Getenv("ZOOKEEPER_URI"),
		ZookeeperTestURI: os.Getenv("ZOOKEEPER_TEST_URI"),
		RedisURI:         os.Getenv("REDIS_URI"),
		RedisPass:        os.Getenv("REDIS_PASS"),
		Port:             os.Getenv("PORT"),
	}
}

// GetEnv ...
func GetEnv() *ENV {
	return &env
}
