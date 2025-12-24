package config

import (
	"fmt"
	"os"
)

type Config struct {
	AppEnv  string
	AppPort string

	DBURL string
}

func Load() (*Config, error) {
	cfg := &Config{
		AppEnv:  getEnv("APP_ENV", "local"),
		AppPort: getEnv("APP_PORT", "8080"),
		DBURL:   os.Getenv("DB_URL"),
	}

	if cfg.DBURL == "" {
		return nil, fmt.Errorf("DB_URL is required but not set")
	}

	return cfg, nil
}

func getEnv(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}
