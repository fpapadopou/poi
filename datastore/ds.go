// Package datastore is an implementation-agnostic data store provider
package datastore

import (
	model "github.com/fpapadopou/poi/datastore/model"
)

// Database interface provides generic database functionality (i.e. connection, etc)
type Database interface {
	Connect(dbConfig model.DatabaseConfig) (*model.Connection, error)
}

// POIManager interface contains POI data insertion/retrieval functionality
type POIManager interface {
	GetAll() ([]*model.POIs, error)
	GetSinglePOIByID(ID model.PrimaryKey) (*model.POIs, error)
}

// DS (DataStore) provides all the database functionality used by the app
type DS struct {
	Connection *model.Connection
	DB         *Database
	POIs       *POIManager
}
