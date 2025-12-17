package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	LinkedInEmail    string
	LinkedInPassword string
}

func Load() *Config {
	// Load .env file if present (safe in production too)
	_ = godotenv.Load()

	cfg := &Config{
		LinkedInEmail:    os.Getenv("LINKEDIN_EMAIL"),
		LinkedInPassword: os.Getenv("LINKEDIN_PASSWORD"),
	}

	if cfg.LinkedInEmail == "" || cfg.LinkedInPassword == "" {
		log.Fatal("LINKEDIN_EMAIL or LINKEDIN_PASSWORD not set")
	}

	return cfg
}
