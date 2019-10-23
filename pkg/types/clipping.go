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
	ID                 uint32     `gorm:"primary_key" json:"id"`
	TitleID            uint32     `json:"-"`
	Title              *Title     `json:"title"`
	LocationRangeStart uint32     `json:"location_start"`
	LocationRangeEnd   uint32     `json:"location_end"`
	PageRangeStart     uint32     `json:"page_start"`
	PageRangeEnd       uint32     `json:"page_end"`
	Type               string     `json:"type"`
	Date               time.Time  `json:"date"`
	Content            string     `json:"content"`
	SourceContent      string     `json:"source_content"`
	Source             string     `json:"-"`
	CreatedAt          time.Time  `json:"-"`
	UpdatedAt          time.Time  `json:"-"`
	DeletedAt          *time.Time `json:"-"`
}
