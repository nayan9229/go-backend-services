package db

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/nayan9229/go-backend-services/chassis"
	"github.com/rs/zerolog/log"
	sqldblogger "github.com/simukti/sqldb-logger"
	"github.com/simukti/sqldb-logger/logadapter/zerologadapter"
)

type MongoClient struct {
	DB *sqlx.DB
}

// NewPGClient creates a new item database connection.
func NewMongoClient(ctx context.Context, dbURL string, debug bool) (*MongoClient, error) {
	db, err := chassis.DBConnectJson(ctx, "item", dbURL)
	if err != nil {
		return nil, err
	}

	if debug {
		db.DB = sqldblogger.OpenDriver(dbURL, db.DB.Driver(), zerologadapter.New(log.Logger))
	}

	return &MongoClient{db}, nil
}
