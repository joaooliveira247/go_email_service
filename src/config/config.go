package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var RabbitURL = ""

func LoadEnv() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatalf("(Config/dotenv): %s\n", err)
	}

	RabbitURL = fmt.Sprintf("amqp://%s:%s@%s:%s/",
		os.Getenv("RABBIT_USER"), os.Getenv("RABBIT_PASSWORD"),
		os.Getenv("RABBIT_HOST"), os.Getenv("RABBIT_PORT"),
	)
}
