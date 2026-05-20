package main

import (
	"fmt"

	"github.com/joaooliveira247/go_email_service/src/config"
)

func init() {
	config.LoadEnv()
}

func main() {
	fmt.Println("Hello Go email service")
}
