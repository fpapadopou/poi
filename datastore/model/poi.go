package model

import (
	"time"
)

// POI struct represents a point of interest object
// This is a shareable model that can be used by all database providers (Postgres, MongoDB, etc)
// The struct also contains metadata that dictates how the object properties should be
// handled during JSON encoding/decoding
type POI struct {
	tableName struct{}   `sql:"poi"`
	ID        PrimaryKey `json:"id,omitempty"`
	Name      string     `json:"name,omitempty"`
	Longitude float64    `json:"lon,omitempty"`
	Latitude  float64    `json:"lat,omitempty"`
	Location  string     `json:"location,omitempty"`
	Category  string     `json:"category,omitempty"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt time.Time  `json:"updated_at,omitempty"`
}
