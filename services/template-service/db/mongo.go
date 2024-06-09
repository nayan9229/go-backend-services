package db

import (
	"context"

	"github.com/nayan9229/go-backend-services/chassis"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoClient struct {
	DB *mongo.Database
}

// NewPGClient creates a new item database connection.
func NewMongoClient(ctx context.Context, dbURL string, debug bool) (*MongoClient, error) {
	db, err := chassis.DBConnectJson(ctx, "genuin", dbURL)
	if err != nil {
		return nil, err
	}
	return &MongoClient{db}, nil
}

func (h *MongoClient) Close() error {
	return h.DB.Client().Disconnect(context.TODO())
}
