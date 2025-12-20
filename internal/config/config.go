package config

import (
	"log"
	"os"

	// Loads environment variables from a .env file
	"github.com/joho/godotenv"
)

// Config holds all application-level configuration values
// These values are kept outside the codebase for security
type Config struct {
	LinkedInEmail    string
	LinkedInPassword string
}

// Load reads configuration values from environment variables
// It ensures required credentials are present before the app starts
func Load() *Config {
	// Load variables from .env file if it exists
	// This is safe for both local development and production
	_ = godotenv.Load()

	// Populate configuration struct from environment variables
	cfg := &Config{
		LinkedInEmail:    os.Getenv("LINKEDIN_EMAIL"),
		LinkedInPassword: os.Getenv("LINKEDIN_PASSWORD"),
	}

	// Validate required configuration
	// Application should not run without credentials
	if cfg.LinkedInEmail == "" || cfg.LinkedInPassword == "" {
		log.Fatal("LINKEDIN_EMAIL or LINKEDIN_PASSWORD not set")
	}

	return cfg
}
