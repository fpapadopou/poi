// +build psql

package model

import (
	"github.com/go-pg/pg"
)

/**
 * This file contains type aliases for properties that are specific to a
 * Postgres database connection.
 **/

// Connection is a type alias for a database connection, in this case
// it aliases a Postgres connection pool
type Connection pg.DB

// PrimaryKey aliases POI primary key to integer, the PK type used by Postgres
type PrimaryKey int

// DatabaseConfig contains database connection info
// The configuration struct definition is compatible with `godotenv` & `env` packages
type DatabaseConfig struct {
	Host     string `env:"DB_HOST" envDefault:"127.0.0.1"`
	Port     string `env:"DB_PORT" envDefault:"5432"`
	Database string `env:"DB_NAME" envDefault:"postgres"`
	User     string `env:"DB_USER" envDefault:"postgres"`
	Password string `env:"DB_PASSWORD" envDefault:""`
}
