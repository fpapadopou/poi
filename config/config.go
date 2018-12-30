package config

import (
	"log"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

// ApplicationConfig contains generic application configuration info
type ApplicationConfig struct {
	Port string `env:"APP_HTTP_PORT" envDefault:"8000"`
}

// DatabaseConfig contains database connection info
type DatabaseConfig struct {
	Host     string `env:"DB_HOST" envDefault:"127.0.0.1"`
	Port     string `env:"DB_PORT" envDefault:"5432"`
	Database string `env:"DB_NAME" envDefault:"postgres"`
	User     string `env:"DB_USER" envDefault:"postgres"`
	Password string `env:"DB_PASSWORD" envDefault:""`
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
func GetDatabaseConfig() DatabaseConfig {
	loadEnv()
	// Parse environment configuration
	var cfg DatabaseConfig
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
