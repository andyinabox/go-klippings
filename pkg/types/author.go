package types

import "time"

// Author encapsulates data for a single author
type Author struct {
	ID         uint32     `json:"id"`
	Name       string     `json:"name"`
	SourceName string     `json:"source_name"`
	Titles     []*Title   `json:"titles,omitempty"`
	CreatedAt  time.Time  `json:"-"`
	UpdatedAt  time.Time  `json:"-"`
	DeletedAt  *time.Time `json:"-"`
}

// AuthorRepository manages authors collection
type AuthorRepository interface {
	Upsert(author *Author) (*Author, error)

	FindOne(ID interface{}) (*Author, error)
	FindAll() ([]*Author, error)

	Search(term string) ([]*Author, error)

	FindForTitle(t *Title) ([]*Author, error)
	FindForClipping(c *Clipping) ([]*Author, error)

	DeleteByID(ID interface{}) error
}
