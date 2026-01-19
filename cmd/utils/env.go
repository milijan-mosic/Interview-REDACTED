package utils

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvVariable(key string) (string, error) {
	value := ""

	err := godotenv.Load()
	if err != nil {
		return value, errors.New("Error loading .env file")
	}

	value = os.Getenv(key)
	if value == "" {
		return value, errors.New("Missing required environment variable: " + key)
	}

	return value, nil
}
