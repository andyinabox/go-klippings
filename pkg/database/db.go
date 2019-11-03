package database

import (
	"github.com/andyinabox/go-klippings-api/pkg/parser"
	"github.com/andyinabox/go-klippings-api/pkg/types"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // has to be in this file
	"log"
	"os"
)

// Database wraps `gorm.DB` with some additional
// application-specific functionality
type Database struct {
	DB   *gorm.DB
	File string
}

// DataImportResult is used to communicate which records
// were imported
type DataImportResult struct {
	Clippings []*types.Clipping
	Authors   []*types.Author
	Titles    []*types.Title
}

// Open opens a database connection to the given
// database file
func Open(fp string) (*Database, error) {
	db, err := gorm.Open("sqlite3", fp)
	if err != nil {
		return nil, err
	}

	// auto-migrate models
	db.AutoMigrate(&types.Clipping{})
	db.AutoMigrate(&types.Title{})
	db.AutoMigrate(&types.Author{})

	wrapper := Database{db, fp}

	return &wrapper, nil
}

// GetAllTitles does exaclty what the name says
func (d *Database) GetAllTitles(t *[]types.Title, deep bool) error {
	if deep {
		d.DB.Preload("Clippings").Preload("Authors").Find(t)
	} else {
		d.DB.Find(t)
	}
	return nil
}

// GetTitlesDeep retrieves titles with all associations
// func (d *Database) GetTitlesDeep(t *[]types.Title) {
// d.DB.Preload("Clippings").Preload("Authors").Find(t)
// }

// GetAuthorsDeep retrieves authors with all associations
// func (d *Database) GetAuthorsDeep(a *[]types.Author) {
// d.DB.Preload("Titles.Clippings").Find(a)
// for _, author := range *a {
// 	for _, t := range author.Titles {
// 		var title types.Title
// 		d.DB.Preload("Clippings").First(&title, t.ID)
// 		t.Clippings = title.Clippings
// 	}
// }
// }

// GetClippingsDeep retrieves clippings with all associations
// func (d *Database) GetClippingsDeep(c *[]types.Clipping) {
// d.DB.Preload("Title.Authors").Find(c)
// for _, clip := range *c {
// 	var title types.Title
// 	d.DB.Preload("Authors").First(&title, clip.Title.ID)
// 	clip.Title.Authors = title.Authors
// }
// }

// Close closes the database connection
func (d *Database) Close() error {
	return d.DB.Close()
}

// Destroy closes the database connection
// **and deletes the database file!**
func (d *Database) Destroy() error {
	var err error
	err = d.DB.Close()
	if err != nil {
		return err
	}
	err = os.Remove(d.File)
	if err != nil {
		return err
	}
	return nil
}

// ProcessParseData takes results from the parser.Parse and
// populates the database, being sure to avoid duplicates
func (d *Database) ProcessParseData(data *[]parser.Data) (*DataImportResult, error) {
	r := &DataImportResult{
		Clippings: make([]*types.Clipping, 0),
		Authors:   make([]*types.Author, 0),
		Titles:    make([]*types.Title, 0),
	}

	for _, p := range *data {
		skip, err := d.ProcessParseDataSingle(&p, r)
		if err != nil {
			return nil, err
		}
		if skip {
			log.Printf("Skipped clipping %v, already exists", p.SourceChecksum)
		}
	}
	return r, nil
}

// ProcessParseDataSingle process an individual parser.Data struct
func (d *Database) ProcessParseDataSingle(p *parser.Data, r *DataImportResult) (bool, error) {

	var cCount int // count matching clippings in db
	var tCount int // count matching titles in db

	d.DB.Model(&types.Clipping{}).Where("id = ?", p.SourceChecksum).Count(&cCount)
	d.DB.Model(&types.Title{}).Where("id = ?", p.TitleChecksum).Count(&tCount)

	// if the clipping already exists we can skip it
	if cCount > 0 {
		return true, nil
	}

	// otherwise we can create a new clipping
	c := types.Clipping{
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
	r.Clippings = append(r.Clippings, &c)

	// Now we'll create or init our title
	var t types.Title
	d.DB.FirstOrInit(&t, &types.Title{
		ID: p.TitleChecksum,
	})
	// if it's new we'll set the title
	// from source
	if t.SourceTitle == "" {
		t.SourceTitle = p.Title
	}
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
		r.Titles = append(r.Titles, &t)
		authors := make([]*types.Author, 0)
		for name, id := range p.Authors {
			a := &types.Author{
				ID:         id,
				Name:       name,
				SourceName: name,
			}
			d.DB.Create(a)
			authors = append(authors, a)
			r.Authors = append(r.Authors, a)
		}

		if len(authors) > 0 {
			d.DB.Model(&t).Association("Authors").Append(&authors)
		}
	}
	return false, nil
}
