package database

import (
	"github.com/andyinabox/go-klippings-api/pkg/parser"
	"os"
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
}
