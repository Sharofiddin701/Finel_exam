package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

const (
	// DebugMode indicates service mode is debug.
	DebugMode = "debug"
	// TestMode indicates service mode is test.
	TestMode = "test"
	// ReleaseMode indicates service mode is release.
	ReleaseMode  = "release"
	SecureApiKey = "1234"
)

type Config struct {
	ServiceName string
	Environment string // debug, test, release
	Version     string

	ServiceHost string
	HTTPPort    string
	HTTPScheme  string
	Domain      string

	DefaultOffset string
	DefaultLimit  string

	UserServiceHost string
	UserGRPCPort    string

	ScheduleServiceHost string
	ScheduleServicePort string
}

// Load ...
func Load() Config {
	if err := godotenv.Load("/go-api-gateway.env"); err != nil {
		fmt.Println("No .env file found")
	}

	config := Config{}

	config.ServiceName = cast.ToString(getOrReturnDefaultValue("SERVICE_NAME", "crm_api_gateway"))
	config.Environment = cast.ToString(getOrReturnDefaultValue("ENVIRONMENT", DebugMode))
	config.Version = cast.ToString(getOrReturnDefaultValue("VERSION", "1.0"))

	config.ServiceHost = cast.ToString(getOrReturnDefaultValue("SERVICE_HOST", "localhost"))
	config.HTTPPort = cast.ToString(getOrReturnDefaultValue("HTTP_PORT", ":8080"))
	config.HTTPScheme = cast.ToString(getOrReturnDefaultValue("HTTP_SCHEME", "http"))
	config.Domain = cast.ToString(getOrReturnDefaultValue("DOMAIN", "localhost:8080"))

	config.DefaultOffset = cast.ToString(getOrReturnDefaultValue("DEFAULT_OFFSET", "0"))
	config.DefaultLimit = cast.ToString(getOrReturnDefaultValue("DEFAULT_LIMIT", "10"))

	config.UserServiceHost = cast.ToString(getOrReturnDefaultValue("USER_SERVICE_HOST", "localhost"))
	config.UserGRPCPort = cast.ToString(getOrReturnDefaultValue("USER_GRPC_PORT", ":9101"))

	config.ScheduleServiceHost = cast.ToString(getOrReturnDefaultValue("SCHEDULE_SERVICE_HOST", "localhost"))
	config.ScheduleServicePort = cast.ToString(getOrReturnDefaultValue("SCHEDULE_SERVICE_PORT", ":9202"))

	return config
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}
