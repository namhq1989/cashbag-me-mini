package config

import(
	"os"
)
type ENV struct {
	IsDev                  bool
	DatabaseURI  string
	Port 	string
}
var env ENV
func initENV(){
	os.Setenv("DatabaseURI", "mongodb+srv://CashbagMe:Cashbag@cluster0.epe8y.gcp.mongodb.net/Cluster0?retryWrites=true&w=majority")
	os.Setenv("Port",":8080")
	env =ENV{
		IsDev : true,
		DatabaseURI: os.Getenv("DatabaseURI"),
		Port :os.Getenv("Port"),
	}
}
func GetEnv() *ENV{
	return &env
}