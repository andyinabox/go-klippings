package types

import ()

// Clipping encapsulates data for a single kindle clipping
type Title struct {
	Title string

	rawTitle string
}

// Create a new clipping
func (c *Title) Create() (Title, error) {
	return Title{}, nil
}
