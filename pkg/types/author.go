package types

import (
	"github.com/jinzhu/gorm"
)

// Clipping encapsulates data for a single kindle clipping
type Author struct {
	gorm.Model
	Name       string
	SourceName string
	ID         int32
	Titles     []*Title `gorm:"many2many:title_authors;"`
}

// Create a new clipping
func (c *Author) Create() (Author, error) {
	return Author{}, nil
}
