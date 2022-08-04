package config

import (
	"os"

	"github.com/spf13/cast"
)

// Config ...
type Config struct {
	Environment string // develop, staging, production

	PostServiceHost string
	PostServicePort int

	DefaultOffset string
	DefaultLimit  string

	CrudServiceHost string
	CrudServicePort int

	Endpoint string

	LogLevel string
	HttpPort string
}

// Load loads environment vars and inflates Config
func Load() Config {
	config := Config{}

	config.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))

	config.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	config.HttpPort = cast.ToString(getOrReturnDefault("HTTP_PORT", ":8090"))

	config.DefaultOffset = cast.ToString(getOrReturnDefault("DEFAULT_OFFSET", "1"))
	config.DefaultLimit = cast.ToString(getOrReturnDefault("DEFAULT_LIMIT", "10"))

	config.PostServiceHost = cast.ToString(getOrReturnDefault("USER_SERVICE_HOST", "localhost"))
	config.PostServicePort = cast.ToInt(getOrReturnDefault("USER_SERVICE_PORT", 5003))

	config.CrudServiceHost = cast.ToString(getOrReturnDefault("CONTACT_SERVICE_HOST", "localhost"))
	config.CrudServicePort = cast.ToInt(getOrReturnDefault("CONTACT_SERVICE_PORT", 5004))
	config.Endpoint = cast.ToString(getOrReturnDefault("ENDPOINT", "https://gorest.co.in"))

	return config
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
