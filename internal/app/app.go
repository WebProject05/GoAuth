package app

import (
	"context"
	"fmt"
	"goauth/internal/config"
	"goauth/internal/db"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	Config config.Config
	MongoClient *mongo.Client
	DB *mongo.Database
}


// Performing the retrival of env variables and connecting to the database
func New(ctx context.Context) (*App, error) {
	cfg, err := config.Load()
	if err != nil {
		return &App{}, fmt.Errorf("Failed to Load and Connect: %w", err)
	}

	mongoCli, err := db.Connect(ctx, cfg)
	if err != nil {
		return &App{}, fmt.Errorf("Failed to connect to database: %w", err)
	}

	return &App{
		Config: cfg,
		MongoClient: mongoCli.Client,
		DB: mongoCli.DB,
	}, nil
}

// Performing Disconnection with the database
func (a *App) Close(ctx context.Context) error {
	if a.MongoClient == nil {
		return nil
	}

	closeCtx, cancel := context.WithTimeout(ctx, 5 * time.Second)
	defer cancel()

	if err := a.MongoClient.Disconnect(closeCtx); err != nil {
		return fmt.Errorf("Mongo Disconnection Failed: %w", err)
	}
	
	return nil
}