package main

import (
	"log"

	"github.com/andyinabox/go-klippings-api/internal/api"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {

	// get settings from `.env` file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// create router
	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.Static("/", "./public")

	// open db
	db, err := gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	// create api
	err = api.Create(router, db)
	if err != nil {
		log.Fatalf("Failed to create API: %v", err)
	}

	// Listen and Server in 0.0.0.0:8080
	router.Run(":8080")
}
