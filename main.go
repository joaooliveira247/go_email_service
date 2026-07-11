package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/joaooliveira247/go_email_service/src/config"
)

func init() {
	handler := slog.NewJSONHandler(
		os.Stdout,
		&slog.HandlerOptions{Level: slog.LevelInfo, AddSource: true},
	)

	logger := slog.New(handler)
	slog.SetDefault(logger)

	config.LoadEnv()
}

func main() {
	slog.Info("Go email service started!")
	fmt.Println("Hello Go email service")
}
