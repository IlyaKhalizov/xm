package config

import (
	"os"
	"time"

	"github.com/joeshaw/envdecode"
	"github.com/rs/zerolog"
)

type Configuration struct {
	Port                int           `env:"HTTP_PORT"`
	HTTPShutdownTimeout time.Duration `env:"HTTP_SHUTDOWN_TIMEOUT,default=5s"`
	HTTPServerTimeout   time.Duration `env:"HTTP_SERVER_TIMEOUT,default=5s"`
	Token               string        `env:"API_TOKEN"`
	Postgres            Postgres
	Redis               Redis
}

type Postgres struct {
	SqlDriver string `env:"SQL_DRIVER"`
	SqlDSN    string `env:"POSTGRESQL_DSN"`
}

type Redis struct {
	Host                string        `env:"REDIS_HOST"`
	Port                int           `env:"REDIS_PORT"`
	DB                  int           `env:"REDIS_DB"`
	CacheExpirationTime int           `env:"CACHE_EXPIRATION_TIME"`
	DefaultCacheTTL     time.Duration `env:"DEFAULT_CACHE_TTL"`
	UseCache            bool          `env:"USE_CACHE"`
}

func Load(logger zerolog.Logger) Configuration {
	var c Configuration
	if err := envdecode.Decode(&c); err != nil {
		logger.Info().Timestamp().Msg(err.Error())
		os.Exit(1)
	}
	return c
}
