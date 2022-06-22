package config

import (
	"os"

	"github.com/spf13/cast"
)

type Config struct {
	Environment      string // develop, staging, production
	PostgresHost     string
	PostgresPort     int
	PostgresDatabase string
	PostgresUser     string
	PostgresPassword string
	LogLevel         string
	HTTPPort         string
	RedisHost        string
	RedisPort        int
	KafkaURL         string
	UserServiceURL   string
}

// Load loads environment vars and inflates Config
func Load() Config {
	c := Config{}

	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))

	c.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "localhost"))
	c.PostgresPort = cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432))
	c.PostgresDatabase = cast.ToString(getOrReturnDefault("POSTGRES_DATABASE", "postgres"))
	c.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "muhammad"))
	c.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "12345"))
	c.RedisHost = cast.ToString(getOrReturnDefault("REDIS_HOST", "localhost"))
	c.RedisPort = cast.ToInt(getOrReturnDefault("REDIS_PORT", 6379))
	c.HTTPPort = cast.ToString(getOrReturnDefault("REDIS_PORT", ":8000"))
	c.KafkaURL = cast.ToString(getOrReturnDefault("KAFKA_URL", "127.0.0.1:9092"))
	c.UserServiceURL = cast.ToString(getOrReturnDefault("USER_SERVICE_URL", "127.0.0.1:9001"))

	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))

	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
