package main

import (
	"log"
	"ushrt/internal/storage"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load("build/.env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	log.Println(".env loaded successfully")
}

func main() {
	database, err := storage.New()
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()
}
