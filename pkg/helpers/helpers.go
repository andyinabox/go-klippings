package helpers

import (
	"github.com/andyinabox/go-klippings-api/internal/database"
	"os"
	"path"
)

const DataDirName = ".klippings"

// func ExecutionDir() (string, error) {
// 	ex, err := os.Executable()
// 	if err != nil {
// 		return "", err
// 	}
// 	exPath := filepath.Dir(ex)
// 	return exPath, nil
// }

// DataDir creates a dir for application data
// and returns the string path
func DataDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	dataDir := path.Join(homeDir, DataDirName)

	// make data dir if it doesn't already exist
	if _, err := os.Stat(dataDir); os.IsNotExist(err) {
		err = os.MkdirAll(dataDir, 0777)
		if err != nil {
			return "", err
		}
	}

	return dataDir, nil
}

// CreateTestDB creates a temp db file for use
// in testing functions
func CreateTestDB() (*database.Database, error) {
	testDb := path.Join(os.TempDir(), "test.db")
	db, err := database.Open(testDb)
	if err != nil {
		return nil, err
	}
	return db, nil
}
