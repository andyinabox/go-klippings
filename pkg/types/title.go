package types

import "time"

// Title encapsulates data for a single title
type Title struct {
	ID          uint32      `gorm:"primary_key" json:"id"`
	Title       string      `json:"title"`
	Clippings   []*Clipping `json:"clippings,omitempty"`
	Authors     []*Author   `gorm:"many2many:title_authors" json:"authors,omitempty"`
	SourceTitle string      `json:"source_title"`
	CreatedAt   time.Time   `json:"-"`
	UpdatedAt   time.Time   `json:"-"`
	DeletedAt   *time.Time  `json:"-"`
}

// TitleRepository is used for managing titles
type TitleRepository interface {
	Upsert(title *Title) (*Title, error)

	FindOne(ID interface{}) (*Title, error)
	FindAll() ([]*Title, error)

	Search(term string) ([]*Title, error)

	FindForAuthor(a *Author) ([]*Title, error)
	FindForClipping(c *Clipping) ([]*Title, error)

	DeleteByID(ID interface{}) error
}
