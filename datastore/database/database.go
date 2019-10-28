package database

import (
	"github.com/fpapadopou/poi/datastore/model"
)

// Provider contains the implementation of several database methods
type Provider interface {
	GetAll() ([]*model.POI, error)
	GetPOIByID(ID model.PrimaryKey) (*model.POI, error)
	CreatePOI(p *model.POI) (*model.POI, error)
	UpdatePOI(p *model.POI) error
	DeletePOI(p *model.POI) error
}

func NewProvider(connection *Connection) Provider {
	return &POIService{
		Conn: connection,
	}
}