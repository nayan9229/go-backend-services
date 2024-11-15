package db

import (
	"context"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/nayan9229/go-backend-services/chassis"
	"github.com/rs/zerolog/log"
	sqldblogger "github.com/simukti/sqldb-logger"
	"github.com/simukti/sqldb-logger/logadapter/zerologadapter"
)

var schema = `CREATE TABLE IF NOT EXISTS users (
	user_id    BIGSERIAL PRIMARY KEY,
    first_name VARCHAR(80)  DEFAULT '',
    last_name  VARCHAR(80)  DEFAULT '',
	email      VARCHAR(250) DEFAULT '',
	password   VARCHAR(250) DEFAULT NULL
);`

type LHClient struct {
	DB *sqlx.DB
}

// NewPGClient creates a new item database connection.
func NewLHClient(ctx context.Context, dbURL string, debug bool) (*LHClient, error) {
	db, err := chassis.DBConnect(ctx, "users", dbURL, "sqlite3", nil, nil)
	if err != nil {
		return nil, err
	}

	if debug {
		db.DB = sqldblogger.OpenDriver(dbURL, db.DB.Driver(), zerologadapter.New(log.Logger))
	}

	// TODO: Replace with migration
	db.MustExec(schema)

	return &LHClient{db}, nil
}

func (lh *LHClient) Close() error {
	return lh.DB.Close()
}
