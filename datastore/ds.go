// Package datastore is an implementation-agnostic data store provider
package datastore

import (
	"github.com/fpapadopou/poi/datastore/database"
)

// DS (DataStore) provides all the database functionality used by the app
type DS struct {
	DatabaseProvider *database.POIService
}

// NewDS creates a new DataStore and returns a reference to it
func NewDS(dbConfig database.Config) *DS {
	connection := database.Connect(dbConfig)
	service := &database.POIService{
		Conn: connection,
	}

	return &DS{DatabaseProvider: service}
}
