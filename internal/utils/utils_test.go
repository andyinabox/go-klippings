package utils

import (
	"os"
	"testing"
)

func TestDataDir(t *testing.T) {
	dataDir, err := DataDir()
	if err != nil {
		t.Fatalf("Error getting data dir: %v", err)
	}

	_, err = os.Stat(dataDir)
	if err != nil {
		t.Fatalf("%v", err)
	}
}
