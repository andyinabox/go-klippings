package types

import ()

// Clipping encapsulates data for a single kindle clipping
type Author struct {
	Name string

	rawName string
}

// Create a new clipping
func (c *Author) Create() (Author, error) {
	return Author{}, nil
}
