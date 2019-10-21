package types

import (
	"github.com/jinzhu/gorm"
	"hash"
	"time"
)

type ClippingType string

const (
	BookmarkClippingType  = ClippingType("Bookmark")
	HighlightClippingType = ClippingType("Highlight")
)

// Clipping encapsulates data for a single kindle clipping
type Clipping struct {
	gorm.Model
	Hash               hash.Hash `gorm:"UNIQUE_INDEX"`
	TitleID            uint
	Title              Title
	LocationRangeStart uint
	LocationRangeEnd   uint
	PageRangeStart     uint
	PageRangeEnd       uint
	Type               ClippingType
	Date               *time.Time
	Content            string
	Raw                string
}

// Create a new clipping
func (c *Clipping) Create() (Clipping, error) {
	return Clipping{}, nil
}

func (c *Clipping) LocationRange() [2]uint {
	return [2]uint{c.LocationRangeStart, c.LocationRangeEnd}
}

func (c *Clipping) PageRange() [2]uint {
	return [2]uint{c.PageRangeStart, c.PageRangeEnd}
}
