package main

import (
	"flag"
	"fmt"
	"os"

	config "github.com/fpapadopou/poi/config"
	"github.com/go-pg/migrations"
	"github.com/go-pg/pg"
)

const usageText = `This program runs "command" on the db. Supported commands are:
  - up [target] - runs all available migrations by default or up to target one if argument is provided.
  - down - reverts last migration.
  - reset - reverts all migrations.
  - version - prints current db version.
  - set_version [version] - sets db version without running migrations.
Usage:
  go run *.go <command> [args]
First time, initialize migrations table: 
  go run migrations/*.go init
`

func main() {
	// Load db config
	dbConfig := config.GetDatabaseConfig()

	// Parse input params
	flag.Usage = usage
	flag.Parse()

	// Connect to Postgres and perform the specified command
	db := pg.Connect(&pg.Options{
		User:     dbConfig.User,
		Database: dbConfig.Database,
		Addr:     dbConfig.Host + ":" + dbConfig.Port,
		Password: dbConfig.Password,
	})

	oldVersion, newVersion, err := migrations.Run(db, flag.Args()...)
	if err != nil {
		exitf(err.Error())
	}
	if newVersion != oldVersion {
		fmt.Printf("Migrated from version %d to %d\n", oldVersion, newVersion)
	} else {
		fmt.Printf("Version is %d\n", oldVersion)
	}
}

func usage() {
	fmt.Printf(usageText)
	flag.PrintDefaults()
	os.Exit(2)
}

func errorf(s string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, s+"\n", args...)
}

func exitf(s string, args ...interface{}) {
	errorf(s, args...)
	os.Exit(1)
}
