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
	ID            int32
	TitleID       int32
	Title         Title
	LocationRange [2]uint32
	PageRange     [2]uint32
	Type          string
	Date          time.Time
	Content       string
	SourceContent string
	Source        string
}
