package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func Config() (string, string, string) {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	mongo_url := os.Getenv("MONGODB_URL")
	secret := os.Getenv("ACCESS_SECRET")
	return port, mongo_url, secret
}
