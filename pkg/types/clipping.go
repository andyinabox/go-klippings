package types

import (
	"github.com/jinzhu/gorm"
	"time"
)

// type ClippingType string

// const (
// 	BookmarkClippingType  = ClippingType("Bookmark")
// 	HighlightClippingType = ClippingType("Highlight")
// )

// Clipping encapsulates data for a single kindle clipping
type Clipping struct {
	gorm.Model
	Checksum           string `gorm:"UNIQUE_INDEX"`
	TitleID            uint
	Title              Title
	LocationRangeStart uint
	LocationRangeEnd   uint
	PageRangeStart     uint
	PageRangeEnd       uint
	Type               string
	Date               time.Time
	Content            string
	RawContent         string
	RawTitle           string
	RawAuthors         string
	Raw                string
}

func (c *Clipping) LocationRange() [2]uint {
	return [2]uint{c.LocationRangeStart, c.LocationRangeEnd}
}

func (c *Clipping) PageRange() [2]uint {
	return [2]uint{c.PageRangeStart, c.PageRangeEnd}
}
