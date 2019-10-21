package utils

import (
	"os"
	"testing"
)

func TestExecutionDir(t *testing.T) {
	path, err := ExecutionDir()
	if err != nil {
		t.Fatalf("Error getting extecution dir: %v", err)
	}
	t.Logf(path)
}

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

func TestLoadDotEnv(t *testing.T) {
	LoadDotEnv()
}

func TestCreateTestDb(t *testing.T) {
	_, err := CreateTestDB()
	if err != nil {
		t.Fatalf("Error creating test db: %v", err)
	}
}
