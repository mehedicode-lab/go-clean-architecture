package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

// Rds holds configuration details for the database connection.
type Rds struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
}

// JwtConfig holds JWT-related configuration values.
type JwtConfig struct {
	AccessSecret  string
	RefreshSecret string
	RefreshTTL    int
	AccessTTL     int
}

// Config holds the overall configuration, including database and JWT settings.
type Config struct {
	Rds       Rds
	JwtConfig JwtConfig
}

// AppConfig is a global variable for accessing the application configuration.
var AppConfig *Config

// LoadConfig loads configuration from the .env file and parses relevant settings.
func LoadConfig() {
	// Load the .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file") // Exit if loading .env fails
	}

	// Parse JWT access and refresh token TTLs (default values if parsing fails)
	accessTTL := parseOrDefaultTTL(os.Getenv("JWT_ACCESS_TTL"), 5*time.Minute)
	refreshTTL := parseOrDefaultTTL(os.Getenv("JWT_REFRESH_TTL"), 30*time.Minute)

	// Initialize AppConfig with the loaded values
	AppConfig = &Config{
		Rds: Rds{
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Name:     os.Getenv("DB_NAME"),
		},
		JwtConfig: JwtConfig{
			AccessSecret:  os.Getenv("JWT_ACCESS_SECRET"),
			RefreshSecret: os.Getenv("JWT_REFRESH_SECRET"),
			RefreshTTL:    int(refreshTTL.Seconds()),
			AccessTTL:     int(accessTTL.Seconds()),
		},
	}
}

// parseOrDefaultTTL parses a TTL value or returns a default if there's an error.
func parseOrDefaultTTL(duration string, defaultTTL time.Duration) time.Duration {
	parsedTTL, err := parseTTL(duration)
	if err != nil {
		log.Printf("Error parsing TTL '%s', using default value: %v", duration, defaultTTL)
		return defaultTTL
	}
	return parsedTTL
}

// parseTTL parses a TTL duration string (e.g., "15m", "7d") into a time.Duration.
// If the string ends with 'd' (for days), it's converted to hours for parsing.
func parseTTL(duration string) (time.Duration, error) {
	// Handle duration strings that might end with 'd' (for days)
	if duration != "" && duration[len(duration)-1] == 'd' {
		// Convert days to hours for time parsing
		return time.ParseDuration(duration[:len(duration)-1] + "h")
	}
	// Otherwise, parse it normally
	return time.ParseDuration(duration)
}
