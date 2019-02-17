// +build psql

package database

import (
	"log"

	"github.com/fpapadopou/poi/datastore/model"
)

// POIService encapsulates functions for data transfer from/to the database
type POIService struct {
	Conn *Connection
}

// GetAll returns all POIs found in Postgres DB
func (p *POIService) GetAll() ([]*model.POI, error) {
	var pois []*model.POI
	err := p.Conn.Model(&pois).Select()
	if err != nil {
		log.Printf("Got an error while fetching POIs: %v", err)
		return nil, err
	}
	return pois, nil
}

// GetPOIByID returns a single POI object specified by its database ID
func (p *POIService) GetPOIByID(ID model.PrimaryKey) (*model.POI, error) {
	poi := &model.POI{ID: ID}
	err := p.Conn.Model(poi).Select()
	if err != nil {
		log.Printf("Got an error while fetching a POI: %v", err)
		return nil, err
	}
	return poi, nil
}
