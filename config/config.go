// config/config.go
package config

import (
	"os"
)

type Config struct {
	Port        string
	MariaDBURI  string
	MongoDBURI  string
	GroqAPIKey  string
	JWTSecret   string
}

func Load() *Config {
	return &Config{
		Port:        getEnv("PORT", "8080"),
		MariaDBURI:  getEnv("MARIADB_URI", "root:password@tcp(localhost:3306)/ai_assistant"),
		MongoDBURI:  getEnv("MONGODB_URI", "mongodb://localhost:27017"),
		GroqAPIKey:  getEnv("GROQ_API_KEY", ""),
		JWTSecret:   getEnv("JWT_SECRET", "secret"),
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}