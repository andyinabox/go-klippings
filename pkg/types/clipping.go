package types

import (
	// "github.com/jinzhu/gorm"
	"time"
)

// type ClippingType string

// const (
// 	BookmarkClippingType  = ClippingType("Bookmark")
// 	HighlightClippingType = ClippingType("Highlight")
// )

// Clipping encapsulates data for a single kindle clipping
type Clipping struct {
	ID                 uint32 `gorm:"primary_key"`
	TitleID            uint32
	Title              Title
	LocationRangeStart uint32
	LocationRangeEnd   uint32
	PageRangeStart     uint32
	PageRangeEnd       uint32
	Type               string
	Date               time.Time
	Content            string
	SourceContent      string
	Source             string
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          *time.Time
}
