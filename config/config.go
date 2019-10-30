package config

import (
	"log"

	"github.com/caarlos0/env"
	"github.com/fpapadopou/poi/datastore/database"
	"github.com/joho/godotenv"
)

// ApplicationConfig contains generic application configuration info
type ApplicationConfig struct {
	Port string `env:"APP_HTTP_PORT" envDefault:"8000"`
}

// GetApplicationConfig returns a struct with database connection info
func GetApplicationConfig() ApplicationConfig {
	loadEnv()
	// Parse environment configuration
	var cfg ApplicationConfig
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("Failed to parse application configuration: %v", err)
	}
	return cfg
}

// GetDatabaseConfig returns a struct with database connection info
func GetDatabaseConfig() database.Config {
	loadEnv()
	// Parse environment configuration
	var cfg database.Config
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("Failed to parse database configuration: %v", err)
	}
	return cfg
}

func loadEnv() {
	// Load .env file config, if any
	// .env file path is always relative to the path from where the go routine is being called
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, loading system-wide env vars..")
	}
}
