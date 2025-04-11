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

	// database, _ := storage.Mock()

	service := service.New(database)
	handler := handler.New(service)

	http.HandleFunc("/", handler.LoadView)
	http.HandleFunc("/api/encode", handler.EncodeURL)
	http.HandleFunc("/api/decode", handler.DecodeUrl)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		log.Fatal("SERVER_PORT environment variable is missing")
	}

	log.Printf("Server is running at: http://localhost:%s", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("Error", err)
	}
}
