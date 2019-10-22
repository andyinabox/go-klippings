package database

import (
	"github.com/andyinabox/go-klippings-api/pkg/parser"
	"github.com/andyinabox/go-klippings-api/pkg/types"
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

const testDb = "../../test/tmp/test.db"
const testDataFile = "../../test/data/my_clippings.txt"

func TestOpen(t *testing.T) {
	db, err := Open(testDb)
	if err != nil {
		t.Fatalf("Error opening db: %v", err)
	}
	err = db.DB.Close()
	if err != nil {
		t.Fatalf("Error closing db: %v", err)
	}
	err = os.Remove(testDb)
	if err != nil {
		t.Logf("Error remove test db: %v", err)
	}
}

func TestProcessParseData(t *testing.T) {
	f, err := os.Open(testDataFile)
	if err != nil {
		t.Fatalf("Error opening clippings file: %v", err)
	}
	defer f.Close()

	data, err := parser.Parse(f)
	if err != nil {
		t.Fatalf("Error parsing data: %v", err)
	}

	db, err := Open(testDb)
	if err != nil {
		t.Fatalf("Error opening db: %v", err)
	}

	result, err := db.ProcessParseData(&data)
	if err != nil {
		t.Fatalf("Error processing parsed data: %v", err)
	}

	// t.Logf("ProcessParseData result: %#v\n", result)

	var titles []types.Title
	err = db.GetAllTitles(&titles, true)
	if err != nil {
		t.Fatalf("Error retrieving titles: %v", err)
	}
	if len(titles) < 1 {
		t.Fatal("No titles returned")
	}
	assert.Equal(t, len(titles), len(result.Titles))

	t.Log("Found titles:")
	for _, title := range titles {

		// gather authors
		aList := title.Authors
		authors := make([]string, len(aList))
		for i, a := range aList {
			authors[i] = a.Name
		}

		// gather clippings
		clippings := title.Clippings

		t.Logf("%s by %s; (%d clippings)\n", title.Title, strings.Join(authors, ", "), len(clippings))
	}

	err = db.DB.Close()
	if err != nil {
		t.Fatalf("Error closing db: %v", err)
	}
	err = os.Remove(testDb)
	if err != nil {
		t.Logf("Error remove test db: %v", err)
	}
}

func TestDuplicates(t *testing.T) {
	f, err := os.Open(testDataFile)
	if err != nil {
		t.Fatalf("Error opening clippings file: %v", err)
	}
	defer f.Close()

	data, err := parser.Parse(f)
	if err != nil {
		t.Fatalf("Error parsing data: %v", err)
	}

	db, err := Open(testDb)
	if err != nil {
		t.Fatalf("Error opening db: %v", err)
	}

	_, err = db.ProcessParseData(&data)
	if err != nil {
		t.Fatalf("Error processing parsed data: %v", err)
	}

	// t.Logf("ProcessParseData result: %#v\n", result)

	var initialClippingsCount int
	var initialTitlesCount int
	var initialAuthorsCount int

	var secondClippingsCount int
	var secondTitlesCount int
	var secondAuthorsCount int

	db.DB.Model(&types.Clipping{}).Count(&initialClippingsCount)
	db.DB.Model(&types.Title{}).Count(&initialTitlesCount)
	db.DB.Model(&types.Title{}).Count(&initialAuthorsCount)

	_, err = db.ProcessParseData(&data)
	if err != nil {
		t.Fatalf("Error processing parsed data a second time: %v", err)
	}

	// t.Logf("ProcessParseData result: %#v\n", result)

	db.DB.Model(&types.Clipping{}).Count(&secondClippingsCount)
	db.DB.Model(&types.Title{}).Count(&secondTitlesCount)
	db.DB.Model(&types.Title{}).Count(&secondAuthorsCount)

	if initialClippingsCount != secondClippingsCount {
		t.Fatalf("Expected %v clippings after second upload, found %v", initialClippingsCount, secondClippingsCount)
	}
	if initialTitlesCount != secondTitlesCount {
		t.Fatalf("Expected %v titles after second upload, found %v", initialTitlesCount, secondTitlesCount)
	}
	if initialAuthorsCount != secondAuthorsCount {
		t.Fatalf("Expected %v authors after second upload, found %v", initialAuthorsCount, secondAuthorsCount)
	}

	err = db.DB.Close()
	if err != nil {
		t.Fatalf("Error closing db: %v", err)
	}
	err = os.Remove(testDb)
	if err != nil {
		t.Logf("Error remove test db: %v", err)
	}
}
