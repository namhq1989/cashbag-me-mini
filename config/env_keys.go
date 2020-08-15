package config

//	"os"

// ENV ...
type ENV struct {
	IsDev bool

	ZookeeperURI     string
	ZookeeperTestURI string

	// App port
	AppPort string

	// Database
	Database struct {
		URI      string
		Name     string
		TestName string
	}

	// Redis
	Redis struct {
		URI  string
		Pass string
	}

	UserString string
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
