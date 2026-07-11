package config

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

var RabbitURL = ""

func LoadEnv() {
	var err error

	if err = godotenv.Load(); err != nil {
		slog.Error("Failed when try load .env file", "error", err)
		os.Exit(1)
	}

	RabbitURL = fmt.Sprintf("amqp://%s:%s@%s:%s/",
		os.Getenv("RABBIT_USER"), os.Getenv("RABBIT_PASSWORD"),
		os.Getenv("RABBIT_HOST"), os.Getenv("RABBIT_PORT"),
	)
}
