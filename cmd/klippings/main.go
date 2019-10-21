package main

import (
	"log"

	"github.com/andyinabox/go-klippings-api/internal/api"
	"github.com/andyinabox/go-klippings-api/internal/database"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
	"path"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

const (
	DataDirName = ".klippings"
	DbFileName  = "klippings.db"
)

var dataDir string
var dbFile string

func main() {

	// get settings from `.env` file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Panicf("Could not find Home dir: %v", err)
	}

	dataDir = path.Join(homeDir, DataDirName)
	dbFile = path.Join(dataDir, DbFileName)

	// make data dir if it doesn't already exist
	if _, err := os.Stat(dataDir); os.IsNotExist(err) {
		err = os.MkdirAll(dataDir, os.ModeDir)
		if err != nil {
			log.Panicf("Could not load data directory: %v", err)
		}
	}

	// create router
	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.Static("/", "./public")

	// open db
	log.Println("Opening database")
	db, err := database.Open(dbFile)
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
