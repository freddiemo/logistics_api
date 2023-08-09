package config

import (
	"log"

	"github.com/joho/godotenv"
)

func GetEnvVariables() map[string]string {
	params, err := godotenv.Read("./config/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return params
}
