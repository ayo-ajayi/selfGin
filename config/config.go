package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

//This is a wrapper that takes the port, dbUrl and auth secret to return them for use on a global scale.
//Set first argument to empty string "" to use default port 8080.
func Config(port string, dbUrl string, secret string) (string, string, string) {
	
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port_ := os.Getenv(port)
	if port_ == "" {
		port_ = "8080"
	}
	dbUrl_ := os.Getenv(dbUrl)
	secret_ := os.Getenv(secret)
	return port_, dbUrl_, secret_
}
