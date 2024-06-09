package mongo

import (
	"context"
	"exoplanetservice/logger"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	UserName string
	Database string
	Host     string
	Password string
}

type Repository struct {
	cfg  *Config
	conn *mongo.Client
}

func (c *Config) GetMongoURI() string {
	return fmt.Sprintf("mongodb://%s/%s", c.Host, c.Database)
}
func NewRepository(ctx context.Context, config *Config) *Repository {
	// Create MongoDB client options
	clientOptions := options.Client().ApplyURI(config.GetMongoURI())

	// If username and password are provided, set them in the options
	if config.UserName != "" && config.Password != "" {
		clientOptions.SetAuth(options.Credential{
			Username: config.UserName,
			Password: config.Password,
		})
	}

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		logger.Panic(ctx, "Error connecting to MongoDB: %v", err)
		return nil
	}

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		logger.Panic(ctx, "Error pinging MongoDB: %v", err)
		return nil
	}
	return &Repository{conn: client, cfg: config}
}
