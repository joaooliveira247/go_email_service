package config

import (
	"log"

	"github.com/joho/godotenv"
)

var RabbitURL = ""

func LoadEnv() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatalf("(Config/dotenv): %s\n", err)
	}
}
