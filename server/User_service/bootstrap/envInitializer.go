package bootstrap

import (
	"log"

	"github.com/joho/godotenv"
)

func Envinitializer() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
