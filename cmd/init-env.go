package main

import (
	"log"

	"github.com/joho/godotenv"
)

func InitializeEnv() {
	godotenv.Load(".env")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
