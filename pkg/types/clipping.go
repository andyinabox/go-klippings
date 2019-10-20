package types

import (
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
	Authors       []*Author
	Title         *Title
	LocationRange [2]uint
	PageRange     [2]uint
	Type          ClippingType
	Date          time.Time
	Content       string

	hash hash.Hash
	raw  string
}

// Create a new clipping
func (c *Clipping) Create() (Clipping, error) {
	return Clipping{}, nil
}
