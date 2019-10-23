package types

import (
	// "github.com/jinzhu/gorm"
	"time"
)

// Author encapsulates data for a single author
type Author struct {
	ID         uint32     `gorm:"primary_key" json:"id"`
	Name       string     `json:"name"`
	SourceName string     `json:"source_name"`
	Titles     []*Title   `gorm:"many2many:title_authors;" json:"titles,omitempty"`
	CreatedAt  time.Time  `json:"-"`
	UpdatedAt  time.Time  `json:"-"`
	DeletedAt  *time.Time `json:"-"`
}

// Create a new clipping
func (c *Author) Create() (Author, error) {
	return Author{}, nil
}
