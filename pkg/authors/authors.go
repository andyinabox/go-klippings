package authors

import (
	"github.com/andyinabox/go-klippings-api/pkg/types"
)

type Repository struct {
	DB types.DBService
}

func NewRepository(
	db types.DBService,
) *Repository {
	return &Repository{
		DB: db,
	}
}

func (r *Repository) Upsert(author *types.Author) (*types.Author, error) {

}
