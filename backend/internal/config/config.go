package config

import (
	"os"
)

type Config struct {
	GoogleCalendarConfig *GoogleCalendarConfig
}

type GoogleCalendarConfig struct {
	CredentialsFile string
}

func LoadConfig() (*Config, error) {
	return &Config{
		GoogleCalendarConfig: &GoogleCalendarConfig{
			CredentialsFile: getEnvOrDefault("GOOGLE_CALENDAR_CREDENTIALS", "internal/config/credentials.json"),
		},
	}, nil
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
