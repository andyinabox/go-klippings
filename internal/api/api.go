package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"

	"github.com/andyinabox/go-klippings-api/internal/database"
	"github.com/andyinabox/go-klippings-api/pkg/types"
)

const routerGroup = "/api"

var db *database.Database
var router *gin.RouterGroup

func Create(r *gin.Engine, d *database.Database) error {

	router = r.Group(routerGroup)
	db = d

	// Ping test
	router.GET("/ping", ping)

	// Clippings
	// router.GET("/clippings/", getClippings)
	// router.POST("/clippings/", uploadClippings)
	// router.GET("/clippings/:id", getClipping)
	// router.PUT("/clippings/:id", updateClipping)

	// Authors
	// router.GET("/authors/", getAuthors)
	// router.GET("/authors/:id", getAuthor)
	// router.PUT("/authors/:id", updateAuthor)

	// Titles
	router.GET("/titles/", getTitles)
	// router.GET("/titles/:id", getTitle)
	// router.PUT("/titles/:id", updateTitle)

	return nil
}

func getTitles(c *gin.Context) {
	var titles []types.Title
	_ = db.GetAllTitles(&titles)

	c.JSON(http.StatusOK, titles)
}

// func getClippings(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "No clippings yet",
// 	})
// }

// func getClipping(c *gin.Context) {
// 	id := c.Params.ByName("id")

// 	c.JSON(http.StatusOK, gin.H{
// 		"message": fmt.Sprintf("No clippings yet: %v", id),
// 	})
// }

// func updateClipping(c *gin.Context) {
// 	id := c.Params.ByName("id")

// 	c.JSON(http.StatusOK, gin.H{
// 		"message": fmt.Sprintf("No clippings yet: %v", id),
// 	})
// }

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func uploadClippings(c *gin.Context) {
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
