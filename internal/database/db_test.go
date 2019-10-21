package database

import (
	"github.com/andyinabox/go-klippings-api/pkg/parser"
	"github.com/andyinabox/go-klippings-api/pkg/types"
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

	err = db.ProcessParseData(&data)
	if err != nil {
		t.Fatalf("Error processing parsed data: %v", err)
	}

	var titles []types.Title
	err = db.GetAllTitles(&titles)
	if err != nil {
		t.Fatalf("Error retrieving titles: %v", err)
	}
	if len(titles) < 1 {
		t.Fatal("No titles returned")
	}

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
