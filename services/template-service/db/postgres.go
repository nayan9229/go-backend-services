package db

import (
	"context"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/nayan9229/go-backend-services/chassis"
	"github.com/rs/zerolog/log"
	sqldblogger "github.com/simukti/sqldb-logger"
	"github.com/simukti/sqldb-logger/logadapter/zerologadapter"
)

type PGClient struct {
	DB *sqlx.DB
}

// NewPGClient creates a new item database connection.
func NewPGClient(ctx context.Context, dbURL string, debug bool) (*PGClient, error) {
	db, err := chassis.DBConnect(ctx, "genuin", dbURL, "postgres", nil, nil)
	if err != nil {
		return nil, err
	}

	if debug {
		db.DB = sqldblogger.OpenDriver(dbURL, db.DB.Driver(), zerologadapter.New(log.Logger))
	}

	// TODO: Replace with migration
	db.MustExec(schema)

	return &PGClient{db}, nil
}

func (pg *PGClient) Close() error {
	return pg.DB.Close()
}
