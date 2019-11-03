package main

import (
	"flag"
	"fmt"
	"github.com/andyinabox/go-klippings-api/pkg/api"
	"github.com/andyinabox/go-klippings-api/pkg/database"
	"github.com/andyinabox/go-klippings-api/pkg/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/zserge/webview"
	"log"
	"os"
	"path"
)

// Version is the current version
const Version = "0.1.0"

// DbFileName file name for database
const DbFileName = "klippings.db"

var dbFile string

func startServer(r *gin.Engine) {
	r.Run(os.Getenv("PORT"))
}

func main() {

	// get settings from `.env` file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file %v", err)
	}

	// setup flags
	launchGuiPtr := flag.Bool("gui", false, "Run app in experimental GUI")
	flag.Parse()

	// get data dir
	dataDir, err := utils.DataDir()
	if err != nil {
		log.Fatalf("Could not get data dir: %v", err)
	}

	// create router
	router := gin.Default()
	// router.MaxMultipartMemory = 8 << 20 // 8 MiB

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
	}))

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

	router.Static("/home", "./web/dist/")

	if *launchGuiPtr {

		go startServer(router)

		log.Println("Launching GUI mode")
		url := fmt.Sprintf("http://localhost%s/home", os.Getenv("PORT"))
		err = webview.Open("Klippings", url, 600, 800, true)
		if err != nil {
			log.Fatalf("Could not launch gui: %v", err)
		}
	} else {
		startServer(router)
	}

}
