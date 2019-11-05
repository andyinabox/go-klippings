package types

// DBFilters defines filters to be used for a db query
type DBFilters struct {
	Fields interface{}
	Sort   string
	Page   int
	Limit  int
}

// DBService is an abstract interface for the main database
type DBService interface {
	Open(config interface{}) error
	Close() error
	Count(collection string, selector interface{}) (int, error)
	FindOne(collection string, selector interface{}, fields interface{}, model interface{}) error
	FindAll(collection string, selector interface{}, model interface{}, filters *DBFilters) error
	Create(collection string, document interface{}) error
	CreateMany(collection string, documents []interface{}) error
	Upsert(collection string, selector interface{}, document interface{}) (interface{}, error)
	UpdateOne(collection string, selector interface{}, update interface{}) error
	UpdateAll(collection string, selector interface{}, update interface{}) error
	RemoveOne(collection string, selector interface{}) error
	RemoveAll(collection string, selector interface{}) error
	EnsureIndex(collection string, indexDef interface{}) error
	DropIndex(collection string, key []string) error
	DropCollection(collection string) error
}
