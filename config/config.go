package config

import (
	"os"

	"go.mongodb.org/mongo-driver/mongo"
)

var Config *AppConfig

// AppConfig holds application-wide configurations
type AppConfig struct {
	Database  string
	DBClient  *mongo.Client
	SecretKey []byte
}

// LoadConfig loads the application configuration from environment variables
func LoadConfig() {
	Config = &AppConfig{
		Database:  "gigmile",
		DBClient:  InitMongo(getEnv("MONGO_DB_URI", "mongodb://localhost:27017")),
		SecretKey: []byte(getEnv("APP_SECRET", "your-secret-key")),
	}
}

// getEnv retrieves the value of an environment variable or returns a default value
func getEnv(key, defaultVal string) string {
	val, exists := os.LookupEnv(key)
	if !exists {
		return defaultVal
	}
	return val
}
