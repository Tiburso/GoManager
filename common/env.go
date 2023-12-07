package common

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvWithDefault(key, defaultValue string) string {
	value := os.Getenv(key)

	if value == "" {
		return defaultValue
	}

	return value
}

func LoadEnv() {
	err := godotenv.Load()

	if err != nil {
		// write a warning to the log
		log.Println("Error loading .env file")
	}
}
