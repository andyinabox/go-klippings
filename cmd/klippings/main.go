package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/andyinabox/go-klippings-api/internal/api"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router, err := api.Create()
	if err != nil {
		log.Fatalf("Failed to create API: %v", err)
	}
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.Static("/", "./public")

	// Listen and Server in 0.0.0.0:8080
	router.Run(":8080")
}
