package config

//	"os"

// ENV ...
type ENV struct {
	IsDev bool

	ZookeeperURI     string
	ZookeeperTestURI string

	// App port
	AppTransactionPort string

	// Database
	Database struct {
		URI            string
		TransactonName string
		TestName       string
	}

	// Redis
	Redis struct {
		URI  string
		Pass string
	}

	// GRPC
	GRPCUri string
}

var env ENV

// InitENV ...
func InitENV() {
	env = ENV{
		IsDev: true,
	}
}

// GetEnv ...
func GetEnv() *ENV {
	return &env
}
