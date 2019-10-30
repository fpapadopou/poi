// +build psql

package model

/**
 * This file contains type aliases for properties that are specific to a
 * Postgres database connection.
 **/

// PrimaryKey aliases POI primary key to integer, the PK type used by Postgres
type PrimaryKey = int
