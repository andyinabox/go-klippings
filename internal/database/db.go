package database

import (
	"github.com/andyinabox/go-klippings-api/pkg/parser"
	. "github.com/andyinabox/go-klippings-api/pkg/types"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
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

	db.AutoMigrate(&Clipping{})
	db.AutoMigrate(&Title{})
	db.AutoMigrate(&Author{})

	return &Database{db}, nil
}

func (db *Database) ProcessParseData(data *[]parser.Data) error {
	for _, d := range *data {
		err := db.ProcessParseDataSingle(&d)
		if err != nil {
			return err
		}
	}
	return nil
}

func (db *Database) ProcessParseDataSingle(d *parser.Data) error {
	log.Println("ProcessParseDataSingle not implemented yet")

	var c Clipping
	var t Title
	var authors []Author

	db.DB.FirstOrInit(&c, Clipping{
		ID: d.SourceChecksum,
	})

	db.DB.FirstOrInit(&t, Title{
		ID: d.TitleChecksum,
	})

	for _, id := range d.Authors {
		var a Author
		db.DB.FirstOrInit(&a, id)
		authors = append(authors, a)
	}

	// c := Clipping{
	// 	ID:                 d.SourceChecksum,
	// 	TitleID:            d.TitleChecksum,
	// 	LocationRangeStart: d.LocationRange[0],
	// 	LocationRangeEnd:   d.LocationRange[1],
	// 	PageRangeStart:     d.PageRange[0],
	// 	PageRangeEnd:       d.PageRange[1],
	// 	Type:               d.Type,
	// 	Date:               d.Date,
	// 	Content:            d.Content,
	// 	SourceContent:      d.Content,
	// 	Source:             d.Source,
	// }

	// var authors = make([]*Author, 0)
	// for name, id := range d.Authors {
	// 	a := &Author{
	// 		ID:         id,
	// 		Name:       name,
	// 		SourceName: name,
	// 	}
	// 	authors = append(authors, a)
	// }

	// t := Title{
	// 	ID:          d.TitleChecksum,
	// 	Title:       d.Title,
	// 	SourceTitle: d.Title,
	// }

	return nil
}

// func CheckTitle

// func CheckAuthor

// func CheckClipping
