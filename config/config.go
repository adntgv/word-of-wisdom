package config

import (
	"os"
)

type Config struct {
	Host string
	Port string
}

func New() (*Config, error) {
	return &Config{
		Host: getEnv("HOST", ""),
		Port: getEnv("PORT", "8080"),
	}, nil
}

func getEnv(key, defaultValue string) string {
	val := os.Getenv(key)

	if val == "" {
		return defaultValue
	}

	return val
}
