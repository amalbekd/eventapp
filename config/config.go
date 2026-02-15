package config

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadConfig() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on system envs")
	}
}