package main

import (
	"log"

	"github.com/andyinabox/go-klippings-api/internal/api"
	"github.com/andyinabox/go-klippings-api/internal/database"
	"github.com/andyinabox/go-klippings-api/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"path"
)

const DbFileName = "klippings.db"

var dbFile string

func main() {

	// get settings from `.env` file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file %v", err)
	}

	// get data dir
	dataDir, err := utils.DataDir()
	if err != nil {
		log.Fatalf("Could not get data dir: %v", err)
	}

	// create router
	router := gin.Default()
	// router.MaxMultipartMemory = 8 << 20 // 8 MiB
	// router.Static("/public/", "./public")

	// open db
	log.Println("Opening database")
	dbFile = path.Join(dataDir, DbFileName)
	db, err := database.Open(dbFile)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	// create api
	log.Println("Setting up API")
	err = api.Create(router, db)
	if err != nil {
		log.Fatalf("Failed to create API: %v", err)
	}

	// Listen and Server in 0.0.0.0:8080
	router.Run(":8080")
}
