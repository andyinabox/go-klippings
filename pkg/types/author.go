package types

import (
	// "github.com/jinzhu/gorm"
	"time"
)

// Clipping encapsulates data for a single kindle clipping
type Author struct {
	ID         uint32 `gorm:"primary_key"`
	Name       string
	SourceName string
	Titles     []*Title `gorm:"many2many:title_authors;"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
}

// Create a new clipping
func (c *Author) Create() (Author, error) {
	return Author{}, nil
}
