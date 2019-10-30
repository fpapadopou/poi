// Package datastore is an implementation-agnostic data store provider
package datastore

import (
	"github.com/fpapadopou/poi/datastore/database"
)

// DS (DataStore) provides all the database functionality used by the app
type DS struct {
	DatabaseProvider database.Provider
}

// NewDS creates a new DataStore and returns a reference to it
func NewDS(dbConfig database.Config) *DS {
	provider := database.NewProvider(database.Connect(dbConfig))
	return &DS{
		DatabaseProvider: provider,
	}
}
