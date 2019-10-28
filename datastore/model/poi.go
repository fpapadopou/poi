package model

import (
	"encoding/json"
	"time"
)

// POI struct represents a point of interest object
// This is a shareable model that can be used by all database providers (Postgres, MongoDB, etc)
// The struct also contains metadata that dictates how the object properties should be
// handled during JSON encoding/decoding as well as Postgres operations
type POI struct {
	tableName struct{}   `sql:"poi"`
	ID        PrimaryKey `json:"id,omitempty",sql:",pk"`
	Name      string     `json:"name,omitempty",sql:",notnull"`
	Longitude float64    `json:"lon,omitempty",sql:",notnull"`
	Latitude  float64    `json:"lat,omitempty",sql:",notnull"`
	Location  string     `json:"location,omitempty"`
	Category  json.RawMessage     `json:"category,omitempty"`
	CreatedAt time.Time  `json:"created_at,omitempty",sql:"default:now()"`
	UpdatedAt time.Time  `json:"updated_at,omitempty"`
	DeletedAt time.Time  `json:"updated_at,omitempty",pg:",soft_delete"`
}
