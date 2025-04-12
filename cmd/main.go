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

	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		log.Fatal("SERVER_PORT environment variable is missing")
	}

	service := service.New(database)
	handler := handler.New(service)

	http.HandleFunc("/", handler.LoadView)
	http.HandleFunc("/r/", handler.Redirect)
	http.HandleFunc("/api/encode", handler.ProcessUrl)

	log.Printf("Server is running at: http://localhost:%s", serverPort)
	if err := http.ListenAndServe(":"+serverPort, nil); err != nil {
		log.Fatal("Error", err)
	}
}
