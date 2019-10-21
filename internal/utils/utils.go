package utils

import (
	"os"
	"path"
)

const DataDirName = ".klippings"

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
