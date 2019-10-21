package types

import (
	"time"
)

// Clipping encapsulates data for a single kindle clipping
type Title struct {
	ID          uint32 `gorm:"primary_key"`
	Title       string
	Clippings   []*Clipping
	Authors     []*Author `gorm:"many2many:title_authors;"`
	SourceTitle string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

// Create a new clipping
func (c *Title) Create() (Title, error) {
	return Title{}, nil
}
