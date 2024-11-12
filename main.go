package main

import (
	"log"
	"main/routes"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	port := os.Getenv("PORT")

	// get router
	router := routes.GetRouter()

	// Log "Server has started" after the server starts listening
	go func() {
		log.Printf("Server has started and is listening on port %s...", port)
	}()

	err = http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
