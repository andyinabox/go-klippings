package types

import (
	"github.com/jinzhu/gorm"
)

// Clipping encapsulates data for a single kindle clipping
type Title struct {
	gorm.Model
	Title     string
	Clippings []Clipping
	Authors   []*Author `gorm:"many2many:title_authors;"`
	RawTitle  string
}

// Create a new clipping
func (c *Title) Create() (Title, error) {
	return Title{}, nil
}
