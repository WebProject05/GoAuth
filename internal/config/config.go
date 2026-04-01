package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURI    string
	MongoDBName string
	JWTSecret   string
}

// loading the .env variables into the os variables
func Load() (Config, error) {
	err := godotenv.Load()
	if err != nil {
		return Config{}, fmt.Errorf("Loading env variables failed...")
	}

	cfg := Config{
		MongoURI: strings.TrimSpace(os.Getenv("MONGO_URI")),
		MongoDBName: strings.TrimSpace(os.Getenv("MONGO_DB_NAME")),
		JWTSecret: strings.TrimSpace(os.Getenv("JWT_SECRET")),
	}


	if cfg.MongoURI == "" {
		return Config{}, fmt.Errorf("Missing Mongo URI")
	}

	if cfg.MongoDBName == "" {
		return Config{}, fmt.Errorf("Missing Mongo DB Name")
	}

	if cfg.JWTSecret == "" {
		return Config{}, fmt.Errorf("Missing JWT Key")
	}

	return cfg, nil
}