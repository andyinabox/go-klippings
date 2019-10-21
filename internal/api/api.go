package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"path/filepath"

	"github.com/andyinabox/go-klippings-api/pkg/types"
)

var db *gorm.DB
var router *gin.Engine

func Create(r *gin.Engine, d *gorm.DB) error {

	router, db = r, d

	db.AutoMigrate(&types.Clipping{})
	db.AutoMigrate(&types.Title{})
	db.AutoMigrate(&types.Author{})

	// Ping test
	router.GET("/ping", ping)

	// Clippings
	router.GET("/clippings/", getClippings)
	router.GET("/clippings/:id", getClipping)
	router.PUT("/clippings/:id", updateClipping)

	// Authors
	// router.GET("/authors/", getAuthors)
	// router.GET("/authors/:id", getAuthor)
	// router.PUT("/authors/:id", updateAuthor)

	// Titles
	// router.GET("/titles/", getTitles)
	// router.GET("/titles/:id", getTitle)
	// router.PUT("/titles/:id", updateTitle)

	// file upload
	router.POST("/upload", uploadFile)

	return nil
}

func getClippings(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "No clippings yet",
	})
}

func getClipping(c *gin.Context) {
	id := c.Params.ByName("id")

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("No clippings yet: %v", id),
	})
}

func updateClipping(c *gin.Context) {
	id := c.Params.ByName("id")

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("No clippings yet: %v", id),
	})
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func uploadFile(c *gin.Context) {
	file, err := c.FormFile("file")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("get form err: %v", err.Error()),
		})
		return
	}

	filename := filepath.Base(file.Filename)
	if err := c.SaveUploadedFile(file, filename); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("upload file err: %v", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("File %s uploaded successfully", file.Filename),
	})
}
