package config

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

var (
	RabbitURL       = ""
	EmailQueue      = ""
	EmailRetryQueue = ""
	EmailDLQQueue   = ""
	EmailRetryTTL   = 1000
)

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func LoadEnv() {
	var err error

	if err = godotenv.Load(); err != nil {
		slog.Error("Failed when try load .env file", "error", err)
		os.Exit(1)
	}

	RabbitURL = fmt.Sprintf("amqp://%s:%s@%s:%s/",
		getEnv("RABBIT_USER", ""), getEnv("RABBIT_PASSWORD", ""),
		getEnv("RABBIT_HOST", ""), getEnv("RABBIT_PORT", ""),
	)
}
