package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	CloudflareAppID string `env:"CLOUDFLARE_APP_ID"`
	CloudflareToken string `env:"CLOUDFLARE_TOKEN"`
}

func LoadConfig() *Config {
	// Try to load .env file but don't fail if it doesn't exist
	_ = godotenv.Load()

	appID := os.Getenv("CLOUDFLARE_APP_ID")
	token := os.Getenv("CLOUDFLARE_TOKEN")

	return &Config{
		CloudflareAppID: appID,
		CloudflareToken: token,
	}
}
