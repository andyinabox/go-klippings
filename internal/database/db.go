package database

import (
	"github.com/andyinabox/go-klippings-api/pkg/parser"
	"github.com/andyinabox/go-klippings-api/pkg/types"
	"github.com/jinzhu/gorm"
	"log"
)

type Database struct {
	DB *gorm.DB
}

func Open(fp string) (*Database, error) {
	db, err := gorm.Open("sqlite3", fp)
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&types.Clipping{})
	db.AutoMigrate(&types.Title{})
	db.AutoMigrate(&types.Author{})

	return &Database{db}, nil
}

func (db *Database) ProcessParseData(data *[]parser.Data) error {
	for i, d := range *data {
		log.Printf("%v %v\n", i, d)
	}
	return nil
}

// func CheckTitle

// func CheckAuthor

// func CheckClipping
