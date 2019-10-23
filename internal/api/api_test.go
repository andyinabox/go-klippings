package api

import (
	"github.com/andyinabox/go-klippings-api/internal/database"
	"github.com/andyinabox/go-klippings-api/internal/utils"
	"github.com/andyinabox/go-klippings-api/pkg/parser"
	"github.com/andyinabox/go-klippings-api/pkg/types"
	"github.com/gin-gonic/gin"
	// "github.com/joho/godotenv"
	// "fmt"
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

// func init() {
// 	err := godotenv.Load("../../.env")
// 	if err != nil {
// 		panic("Error loading .env file")
// 	}
// }

const (
	port = ":8080"
)

func createAPI() (*gin.Engine, *database.Database, error) {
	router := gin.Default()

	db, err := utils.CreateTestDB()
	if err != nil {
		return nil, nil, err
	}

	err = Create(router, db)
	if err != nil {
		return nil, nil, err
	}
	return router, db, nil
}

func TestCreateAPI(t *testing.T) {
	_, db, err := createAPI()
	defer db.Destroy()

	if err != nil {
		t.Fatalf("Error creating API: %v", err)
	}
}

func TestRootPath(t *testing.T) {
	router, db, err := createAPI()
	defer db.Destroy()

	assert.Nil(t, err)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	t.Log(w.Body.String())
}

func TestBasicTypes(t *testing.T) {
	router, db, err := createAPI()
	defer db.Destroy()
	assert.Nil(t, err)

	f, err := os.Open("../../test/data/my_clippings.txt")
	assert.Nil(t, err)

	data, err := parser.Parse(f)
	assert.Nil(t, err)

	_, err = db.ProcessParseData(&data)
	assert.Nil(t, err)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/api/clippings", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	var clippings []types.Clipping
	b, err := ioutil.ReadAll(w.Body)
	assert.Nil(t, err)
	err = json.Unmarshal(b, &clippings)
	assert.Nil(t, err)
	assert.Equal(t, 23, len(clippings))

	t.Log("Clippings:")
	for i, o := range clippings {
		t.Logf("%d: %d\n", i, o.ID)
	}

	req, _ = http.NewRequest("GET", "/api/titles", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	var titles []types.Title
	b, err = ioutil.ReadAll(w.Body)
	assert.Nil(t, err)
	err = json.Unmarshal(b, &titles)
	assert.Nil(t, err)
	assert.Equal(t, 3, len(titles))

	t.Log("Titles:")
	for i, o := range titles {
		t.Logf("%d: %s\n", i, o.Title)
	}

	req, _ = http.NewRequest("GET", "/api/authors", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	var authors []types.Author
	b, err = ioutil.ReadAll(w.Body)
	assert.Nil(t, err)
	err = json.Unmarshal(b, &authors)
	assert.Nil(t, err)
	assert.Equal(t, 4, len(authors))

	t.Log("Authors:")
	for i, o := range authors {
		t.Logf("%d: %s\n", i, o.Name)
	}
}

func TestUpload(t *testing.T) {
	router, db, err := createAPI()
	defer db.Destroy()
	assert.Nil(t, err)

	w := httptest.NewRecorder()

	f, err := os.Open("../../test/data/my_clippings.txt")
	assert.Nil(t, err)

	var b bytes.Buffer
	mw := multipart.NewWriter(&b)
	fw, err := mw.CreateFormFile("file", "../../test/data/my_clippings.txt")
	assert.Nil(t, err)

	io.Copy(fw, f)
	mw.Close()

	req, err := http.NewRequest("POST", "/api/clippings", &b)
	assert.Nil(t, err)

	req.Header.Set("Content-Type", mw.FormDataContentType())

	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	j, err := ioutil.ReadAll(w.Result().Body)
	assert.Nil(t, err)

	// var result database.DataImportResult
	// err = json.Unmarshal(j, &result)
	// assert.Nil(t, err)

	t.Logf("Result: %#v", string(j))

}
