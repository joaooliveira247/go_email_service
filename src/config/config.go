package config

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	RabbitURL       = ""
	RabbitExchange   = ""
	EmailQueue      = ""
	EmailRetryQueue = ""
	EmailDLQQueue   = ""
	EmailRetryTTL   = 1000
	MaxRetries      = 3
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

	EmailQueue = getEnv("RABBIT_EMAIL_QUEUE", "email.orders")
	EmailRetryQueue = getEnv("RABBIT_EMAIL_RETRY_QUEUE", "email.orders")
	EmailDLQQueue = getEnv("RABBIT_EMAIL_RETRY_QUEUE", "email.orders")

	if ttlStr := os.Getenv("RABBIT_RETRY_TTL"); ttlStr != "" {
		if val, err := strconv.Atoi(ttlStr); err == nil {
			EmailRetryTTL = val
		} else {
			slog.Warn(
				"RABBIT_RETRY_TTL invalid, change to default value",
				"error",
				err,
			)
		}
	}
}
