package types

import "io"

// ParserData is the raw data format for parsed
// kindle clipping data
type ParserData struct {
	Title         string
	Authors       map[string]uint32
	Content       string
	LocationRange [2]uint32
	PageRange     [2]uint32
	Type          string
	Date          string
	Source        string
}

// ParserService parses data from a kindle clippings fole
type ParserService interface {
	Parse(r io.Reader) ([]ParserData, error)
}
