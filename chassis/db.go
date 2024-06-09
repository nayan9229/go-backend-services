package chassis

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type assetFunc func(name string) ([]byte, error)
type assetDirFunc func(name string) ([]string, error)

// DBConnect creates a new database connection.
func DBConnect(ctx context.Context, dbName string, dbURL string, dbDriver string,
	asset assetFunc, assetDir assetDirFunc) (*sqlx.DB, error) {

	db, err := sqlx.Open(dbDriver, dbURL)
	if err != nil {
		return nil, errors.Wrap(err, "opening "+dbName+" database")
	}
	err = db.Ping()
	if err != nil {
		return nil, errors.Wrap(err, "pinging "+dbName+" database")
	}

	// Limit maximum connections (default is unlimited).
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(2)
	return db, nil
}

func DBConnectJson(ctx context.Context, dbName string, dbURL string) (*mongo.Database, error) {
	loggerOptions := options.
		Logger().
		SetComponentLevel(options.LogComponentCommand, options.LogLevelDebug)

	client, err := mongo.Connect(ctx, options.Client().
		ApplyURI(dbURL).SetLoggerOptions(loggerOptions))
	if err != nil {
		return nil, errors.Wrap(err, "opening "+dbName+" database")
	}
	db := client.Database(dbName)
	err = db.Client().Ping(ctx, nil)
	if err != nil {
		return nil, errors.Wrap(err, "pinging "+dbName+" database")
	}
	return db, nil
}
