package main

import (
	"log"
	"net/http"
	"os"

	"orderpackscalculator/internal/api"
	"orderpackscalculator/internal/config"
)

func main() {
	config.LoadDefaultPackSizes()
	router := api.NewRouter()

	// Get port from env or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server started on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
