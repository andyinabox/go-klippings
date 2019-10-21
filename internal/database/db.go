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

func (d *Database) GetAllTitles(t *[]Title) error {
	d.DB.Preload("Clippings").Preload("Authors").Find(t)
	return nil
}

func (d *Database) ProcessParseData(data *[]parser.Data) error {
	for _, p := range *data {
		skip, err := d.ProcessParseDataSingle(&p)
		if err != nil {
			return err
		}
		if skip {
			log.Printf("Skipped clipping %v, already exists", p.SourceChecksum)
		}
	}
	return nil
}

func (d *Database) ProcessParseDataSingle(p *parser.Data) (bool, error) {

	var cCount int // count matching clippings in db
	var tCount int // count matching titles in db

	d.DB.Model(&Clipping{}).Where("id = ?", p.SourceChecksum).Count(&cCount)
	d.DB.Model(&Title{}).Where("id = ?", p.TitleChecksum).Count(&tCount)

	// if the clipping already exists we can skip it
	if cCount > 0 {
		return true, nil
	}

	// otherwise we can create a new clipping
	c := Clipping{
		ID:                 p.SourceChecksum,
		LocationRangeStart: p.LocationRange[0],
		LocationRangeEnd:   p.LocationRange[1],
		PageRangeStart:     p.PageRange[0],
		PageRangeEnd:       p.PageRange[1],
		Type:               p.Type,
		Date:               p.Date,
		Content:            p.Content,
		SourceContent:      p.Content,
		Source:             p.Source,
	}
	d.DB.Create(&c)

	// Now we'll create or init our title
	var t Title
	d.DB.FirstOrInit(&t, Title{
		ID:          p.TitleChecksum,
		SourceTitle: p.Title,
	})
	// if it's new we'll set the title
	// from source
	if t.Title == "" {
		t.Title = p.Title
	}
	// add the clipping to
	d.DB.Model(&t).Association("Clippings").Append(&c)
	d.DB.Save(&t)

	// this means it's a new title
	// we should really only need to add
	// authors once per title
	if tCount == 0 {
		authors := make([]Author, 0)
		for name, id := range p.Authors {
			a := Author{
				ID:         id,
				Name:       name,
				SourceName: name,
			}
			d.DB.Create(&a)
			authors = append(authors, a)
		}

		if len(authors) > 0 {
			d.DB.Model(&t).Association("Authors").Append(&authors)
		}

	}
	return false, nil
}

// func CheckTitle

// func CheckAuthor

// func CheckClipping
