package cmd

import (
	"log"

	"github.com/joho/godotenv"
)

func Environment() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Not expected an error to load the .env file.")
	}
}
