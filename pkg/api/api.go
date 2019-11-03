package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	// "path/filepath"
	"github.com/andyinabox/go-klippings-api/pkg/database"
	"github.com/andyinabox/go-klippings-api/pkg/parser"
	"github.com/andyinabox/go-klippings-api/pkg/types"
	// "github.com/jinzhu/gorm"
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

	router.GET("/authors/:id", getAuthorsSingle)
	router.PUT("/authors/:id", noop)
	router.DELETE("/authors/:id", noop)

	// Titles
	router.GET("/titles", getTitles)

	router.GET("/titles/:id", getTitlesSingle)
	router.PUT("/titles/:id", noop)
	router.DELETE("/titles/:id", noop)

	router.GET("/search", noop)

	return nil
}

// SingleQueryResponse is a response formay
// used in single record queries
type SingleQueryResponse struct {
	Model     interface{}       `json:"model"`
	Clippings []*types.Clipping `json:"clippings"`
}

func getClippings(c *gin.Context) {
	var o types.Clipping
	getObjectCollection(c, &o)
}

func getAuthors(c *gin.Context) {
	var o types.Author
	getObjectCollection(c, &o)
}

func getTitles(c *gin.Context) {
	var o types.Title
	getObjectCollection(c, &o)
}

func getObjectCollection(c *gin.Context, o interface{}) {

}

func getObject(o interface{}) func() {
	return func(c *gin.Context) {
		// build and execute query from params

		// if there are errors, return 500

		// if no record is found, return 404

		// otherwise return json
	}
}

func putObject(o interface{}) func() {
	return func(c *gin.Context) {
		// build and execute query

		// if there are errors, return 500
	}
}

func deleteObject(o interface{}) func() {
	return func(c *gin.Context) {
		// build and execute query based on params

		// if there are errors, return 500
	}
}

// 404 response
func notFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, H{})
}

// 500 response
func serverError(e error) {
	c.JSON(http.StatusInternalServerError, e)
}

func getRoot(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Nothin to see here",
	})
}
func noop(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "This section not yet implemented",
	})
}

// if err := db.Table("employee").Select("department.id, employee.department_id, employeeContact.employee_id").Joins("JOIN department on department.id = employee.department_id").Joins("JOIN employeeContact on employeeContact.id = employee.id").Find(&results).Error; err != nil {
//     return err, ""
// }

func getClippings(c *gin.Context) {
	var clippings []types.Clipping
	db.DB.Preload("Title.Authors").Find(&clippings)
	c.JSON(http.StatusOK, clippings)
}

func getTitles(c *gin.Context) {
	var titles []types.Title
	db.DB.Preload("Clippings").Preload("Authors").Find(&titles)
	c.JSON(http.StatusOK, titles)
}

func getTitlesSingle(c *gin.Context) {
	var title types.Title
	var clippings []*types.Clipping
	db.DB.Preload("Clippings.Title.Authors").Preload("Authors").First(&title, c.Param("id"))

	clippings = title.Clippings
	title.Clippings = nil

	// fmt.Printf("%#v", clippings)
	c.JSON(http.StatusOK, &SingleQueryResponse{
		Model:     &title,
		Clippings: clippings,
	})
}

func getAuthors(c *gin.Context) {
	var authors []types.Author
	db.DB.Preload("Titles.Clippings").Find(&authors)
	c.JSON(http.StatusOK, authors)
}

func getAuthorsSingle(c *gin.Context) {
	var author types.Author
	clippings := make([]*types.Clipping, 0)
	db.DB.Preload("Titles.Clippings.Title.Authors").First(&author, c.Param("id"))
	for _, t := range author.Titles {
		clippings = append(clippings, t.Clippings...)
		t.Clippings = nil
	}
	// fmt.Printf("%#v", clippings)
	c.JSON(http.StatusOK, &SingleQueryResponse{
		Model:     &author,
		Clippings: clippings,
	})
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
