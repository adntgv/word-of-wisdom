package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Host       string
	Port       string
	Difficulty int
}

func New() (*Config, error) {
	difficulty, err := strconv.Atoi(getEnv("DIFFICULTY", "3"))
	if err != nil {
		return nil, fmt.Errorf("invalid difficulty parameter: %v", err)
	}

	return &Config{
		Host:       getEnv("HOST", ""),
		Port:       getEnv("PORT", "8080"),
		Difficulty: difficulty,
	}, nil
}

func getEnv(key, defaultValue string) string {
	val := os.Getenv(key)

	if val == "" {
		return defaultValue
	}

	return val
}
