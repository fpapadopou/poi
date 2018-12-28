package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/caarlos0/env"
	"github.com/go-pg/migrations"
	"github.com/go-pg/pg"
	"github.com/joho/godotenv"
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

type config struct {
	DatabaseHost string `env:"DB_HOST" envDefault:"127.0.0.1"`
	DatabasePort string `env:"DB_PORT" envDefault:"5432"`
}

func main() {
	// Load .env file config, if any
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, loading system-wide env vars..")
	}
	// Parse environment configuration
	var cfg config
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("Failed to parse environment configuration: %v", err)
	}

	// Parse input params
	flag.Usage = usage
	flag.Parse()

	// Connect to Postgres and perform the specified command
	db := pg.Connect(&pg.Options{
		User:     "postgres",
		Database: "postgres",
		Addr:     cfg.DatabaseHost + ":" + cfg.DatabasePort,
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
