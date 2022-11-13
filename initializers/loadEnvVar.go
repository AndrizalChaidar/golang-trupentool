package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVar() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Cannto load dotenv")
	}
}