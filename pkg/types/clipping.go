package types

import (
	"time"
	"io"
)

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

// ClippingRepository manages the clipping collection
type ClippingRepository interface {
	Upsert(clipping *Clipping) (*Clipping, error)

	FindOne(ID interface{}) (*Clipping, error)
	FindAll() ([]*Clipping, error)

	Search(term string) ([]*Clipping, error)

	FindForTitle(t *Title) ([]*Clipping, error)
	FindForAuthor(a *Author) ([]*Clipping, error)

	DeleteByID(ID interface{}) error
}



// ClippingUploadResult contains information about a clippings
// file upload
type ClippingUploadResult struct {
	Clippings []*Clipping
	Authors   []*Author
	Titles    []*Title
	Errors    []*error
}


// ClippingService handles clipping use cases
type ClippingService interface {
	HandleUpload(r io.Reader) (ClippingUploadResult, error)
	GetClippingsRepo() ClippingRepository
	GetAuthorsRepo() AuthorRepository
	GetTitlesRepo() TitleRepository
}