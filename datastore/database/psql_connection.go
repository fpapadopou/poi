// +build psql

package database

import (
	"github.com/go-pg/pg"
)

// Connect returns a Postgres connection pool handle
// This handle is safe for concurrent connections
func Connect(dbConfig Config) *Connection {
	return pg.Connect(&pg.Options{
		User:     dbConfig.User,
		Database: dbConfig.Database,
		Addr:     dbConfig.Host + ":" + dbConfig.Port,
		Password: dbConfig.Password,
	})
}
