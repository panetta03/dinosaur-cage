package database

import (
	"github.com/hashicorp/go-memdb"
)

var db *memdb.MemDB

// InitDB initializes the go-memdb database.
func InitDB() {
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"dinosaurs": {
				Name: "dinosaurs",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.IntFieldIndex{Field: "ID"},
					},
				},
			},
			// Define more tables and indexes if needed.
		},
	}

	var err error
	db, err = memdb.NewMemDB(schema)
	if err != nil {
		panic(err)
	}
}

// GetDB returns the go-memdb database instance.
func GetDB() *memdb.MemDB {
	return db
}
