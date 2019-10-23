package types

import (
	"time"
)

// Title encapsulates data for a single title
type Title struct {
	ID          uint32      `gorm:"primary_key" json:"id"`
	Title       string      `json:"title"`
	Clippings   []*Clipping `json:"clippings,omitempty"`
	Authors     []*Author   `gorm:"many2many:title_authors" json:"authors,omitempty"`
	SourceTitle string      `json:"source_title"`
	CreatedAt   time.Time   `json:"-"`
	UpdatedAt   time.Time   `json:"-"`
	DeletedAt   *time.Time  `json:"-"`
}

// Create a new clipping
func (c *Title) Create() (Title, error) {
	return Title{}, nil
}
