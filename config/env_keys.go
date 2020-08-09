package config

import(
	"os"
)
type ENV struct {
	IsDev                  bool
	ZookeeperURI           string
	DatabaseURI  string
	RedisURI string
	RedisPass string
	Port 	string
	
}
var env ENV
func initENV(){
	os.Setenv("DatabaseURI", "mongodb+srv://CashbagMe:Cashbag@cluster0.epe8y.gcp.mongodb.net/Cluster0?retryWrites=true&w=majority")
	os.Setenv("ZOOKEEPER_URI","127.0.0.1:2181")
	os.Setenv("RedisURI","localhost:6379")
	os.Setenv("RedisPass","")

	os.Setenv("Port",":8080")
	env =ENV{
		IsDev : true,
		DatabaseURI: os.Getenv("DatabaseURI"),
		ZookeeperURI:           os.Getenv("ZOOKEEPER_URI"),
		RedisURI: os.Getenv("RedisURI"),
		RedisPass: os.Getenv("RedisPass"),
		Port :os.Getenv("Port"),
	}
}
func GetEnv() *ENV{
	return &env
}