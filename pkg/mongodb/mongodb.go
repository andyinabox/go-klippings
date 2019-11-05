package mongodb

import (
	// https://github.com/mongodb/mongo-go-driver
	"go.mongodb.org/mongo-driver/mongo"
)

type DB struct {
	Client mongo.Client
	URI    string
}

func NewDB() *DB {
	return &DB{}
}

func (db *DB) Open() error {
	if db.Client == nil {
		// connect to mongodb, set Client
	}

	return nil
}

func (db *DB) Close() error {

}
