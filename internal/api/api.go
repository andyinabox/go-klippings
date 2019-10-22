package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	// "path/filepath"

	"github.com/andyinabox/go-klippings-api/internal/database"
	"github.com/andyinabox/go-klippings-api/pkg/parser"
	"github.com/andyinabox/go-klippings-api/pkg/types"
)

const routerGroup = "/api"

var db *database.Database
var router *gin.RouterGroup

// Create create API router group
func Create(r *gin.Engine, d *database.Database) error {

	router = r.Group(routerGroup)
	db = d

	router.GET("", getRoot)

	// Clippings
	router.GET("/clippings", getClippings)
	router.POST("/clippings", uploadClippings)

	// Authors
	router.GET("/authors", getAuthors)

	// Titles
	router.GET("/titles", getTitles)

	return nil
}

func getRoot(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Nothin to see here",
	})
}

func getClippings(c *gin.Context) {
	var clippings []types.Clipping
	db.DB.Find(&clippings)
	c.JSON(http.StatusOK, clippings)
}

func getTitles(c *gin.Context) {
	var titles []types.Title
	db.DB.Find(&titles)
	c.JSON(http.StatusOK, titles)
}

func getAuthors(c *gin.Context) {
	var authors []types.Author
	db.DB.Find(&authors)
	c.JSON(http.StatusOK, authors)
}

func uploadClippings(c *gin.Context) {
	fh, err := c.FormFile("file")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("form err: %v", err.Error()),
		})
		return
	}

	f, err := fh.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("form err: %v", err.Error()),
		})
		return
	}

	data, err := parser.Parse(f)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("parse err: %v", err.Error()),
		})
		return
	}

	result, err := db.ProcessParseData(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("db err: %v", err.Error()),
		})
		return
	}
	// filename := filepath.Base(file.Filename)
	// if err := c.SaveUploadedFile(file, filename); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"message": fmt.Sprintf("upload file err: %v", err.Error()),
	// 	})
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("File %s uploaded successfully", fh.Filename),
		"records": result,
	})
}
