// Package datastore is an implementation-agnostic data store provider
package datastore

import (
	model "github.com/fpapadopou/poi/datastore/model"
)

// POIService interface contains POI data insertion/retrieval functionality
type POIService interface {
	GetAll() ([]*model.POIs, error)
	GetSinglePOIByID(ID model.PrimaryKey) (*model.POIs, error)
}

// DS (DataStore) provides all the database functionality used by the app
type DS struct {
	Connection *model.Connection
	POIs       *POIService
}
