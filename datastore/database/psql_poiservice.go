// +build psql

package database

import (
	"log"

	"github.com/fpapadopou/poi/datastore/model"
	"github.com/go-pg/pg"
)

// POIService encapsulates functions for data transfer from/to the database
type POIService struct {
	Conn *Connection
}

// TODO: Replace errors when returning with Errors.new()
// GetAll returns all POIs found in Postgres DB.
func (p *POIService) GetAll() ([]*model.POI, error) {
	var pois []*model.POI
	err := p.Conn.Model(&pois).Select()
	if err != nil {
		log.Printf("Got an error while fetching POIs: %v", err)
		return nil, err
	}
	return pois, nil
}

// GetPOIByID returns a single POI object specified by its database ID.
func (p *POIService) GetPOIByID(ID model.PrimaryKey) (*model.POI, error) {
	poi := &model.POI{ID: ID}
	err := p.Conn.Model(poi).Select()
	if err != nil {
		log.Printf("Got an error while fetching a POI: %v", err)
		return nil, err
	}
	return poi, nil
}

// CreatePOI adds a new POI to the store, creating a POSTGis point for the given lat/lon of the POI.
func (p *POIService) CreatePOI(poi *model.POI) (*model.POI, error) {
	// Since Postgres does not return the last insert ID by default, `RETURNING` keyword comes in handy
	_, err := p.Conn.Model(&model.POI{}).Query(
		pg.Scan(poi.ID, poi.Location, poi.CreatedAt),
		`INSERT INTO poi (name, longitude, latitude, category, location) 
		(?name, ?longitude, ?latitude, ?category, ST_MakePoint(?longitude, ?latitude))
		RETURNING id, location, created_at`,
		poi,
	)

	if err != nil {
		log.Printf("Got an error while creating POI: %v", err)
		return nil, err
	}

	return poi, nil
}

// UpdatePOI updates the provided POI in the database and returns an error (if any).
func (p *POIService) UpdatePOI(poi *model.POI) error {
	return nil
}

// DeletePOI will return a POI object from the store and return an error (if any).
func (p *POIService) DeletePOI(poi *model.POI) error {

	return nil
}
