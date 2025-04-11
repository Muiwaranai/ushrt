package main

import (
	"log"
	"net/http"
	"os"
	"ushrt/internal/handler"
	"ushrt/internal/service"
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

	service := service.New(database)
	handler := handler.New(service)

	http.HandleFunc("/", handler.LoadView)

	log.Println("Server running on port: ", os.Getenv("SERVER_PORT"))
	if err := http.ListenAndServe(":"+os.Getenv("SERVER_PORT"), nil); err != nil {
		log.Fatal("Error", err)
	}
}
