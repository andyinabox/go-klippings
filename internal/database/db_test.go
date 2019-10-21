package database

import (
	"os"
	"testing"
)

func TestOpen(t *testing.T) {
	testDb := "../../test/tmp/test.db"
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
