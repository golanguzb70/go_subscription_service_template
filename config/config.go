package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	Environment      string // develop, staging, production
	RPCPort          string
	PostgresHost     string
	PostgresPort     int
	PostgresDatabase string
	PostgresUser     string
	PostgresPassword string
	CloudURL         string
}

func Load() *Config {
	c := &Config{}
	godotenv.Load(cast.ToString(getOrReturnDefault("DOT_ENV_PATH", ".env")))

	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))
	c.RPCPort = cast.ToString(getOrReturnDefault("GRPC_PORT", ":9000"))

	c.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "localhost"))
	c.PostgresPort = cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432))
	c.PostgresDatabase = cast.ToString(getOrReturnDefault("POSTGRES_DATABASE", "subscription_service"))
	c.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "subscription_service"))
	c.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "umeechai3wei3yoh"))

	c.CloudURL = cast.ToString(getOrReturnDefault("CLOUD_URL", "https://test.cdn.uzdigital.tv/uzdigital/images/"))

	return c
}

func getOrReturnDefault(key string, defaultValue any) any {
	v, exists := os.LookupEnv(key)
	if exists {
		return v
	}

	return defaultValue
}
