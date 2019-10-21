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

func (db *Database) ProcessParseDataSingle(d *parser.Data) (bool, error) {

	// var c Clipping
	// var t Title
	// var authors []Author

	var count int
	db.DB.Model(&Clipping{}).Where("id = ?", d.SourceChecksum).Count(&count)

	// if the clipping already exists we can skip it
	if count > 0 {
		return false, nil
	}

	// db.DB.First(&c, d.SourceChecksum)

	log.Printf("%v records found\n", count)

	// for name, id := range d.Authors {
	// 	var a Author
	// 	db.DB.FirstOrInit(&a, id)
	// 	a.Name = name
	// 	a.SourceName = name
	// 	authors = append(authors, a)
	// }

	// db.DB.FirstOrInit(&t, Title{
	// 	ID: d.TitleChecksum,
	// })

	// db.Model(&Title).Association("Clippings").Apppend(&c)

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

	return true, nil
}

// func CheckTitle

// func CheckAuthor

// func CheckClipping
