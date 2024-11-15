package server

import (
	"context"
	"time"

	"github.com/nayan9229/go-backend-services/chassis"
	"github.com/nayan9229/go-backend-services/services/template-service/db"
	"github.com/rs/zerolog/log"
)

type Server struct {
	chassis.Server
	stageEnv bool
	sqlDb    db.DB
	jsonDb   db.JsonDB
}

type Config struct {
	chassis.ServerConfig
	DevMode    bool   `env:"DEV_MODE,default=false"`
	DBQueryLog bool   `env:"DATABASE_QUERY_LOG,default=true"`
	DBURL      string `env:"DATABASE_URL,require=true"`
	JSON_DBURL string `env:"JSON_DATABASE_URL,require=true"`
}

func NewServer(cfg *Config) *Server {
	s := &Server{
		stageEnv: cfg.DevMode,
	}
	s.Init(&cfg.ServerConfig, s.routes())
	var err error

	timeout, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	// s.sqlDb, err = db.NewPGClient(timeout, cfg.DBURL, cfg.DBQueryLog)
	// if err != nil {
	// 	log.Fatal().Err(err).Msg("couldn't connect to user database")
	// }

	s.sqlDb, err = db.NewLHClient(timeout, cfg.DBURL, cfg.DBQueryLog)
	if err != nil {
		log.Fatal().Err(err).Msg("couldn't connect to user database")
	}

	// s.jsonDb, err = db.NewMongoClient(timeout, cfg.JSON_DBURL, cfg.DBQueryLog)
	// if err != nil {
	// 	log.Fatal().Err(err).Msg("couldn't connect to user database")
	// }

	s.AddAtExit(func() {
		if err := s.sqlDb.Close(); err != nil {
			log.Fatal().Err(err).Msg("couldn't disconnect to mongo database")
		}
	})

	// s.AddAtExit(func() {
	// 	if err := s.jsonDb.Close(); err != nil {
	// 		log.Fatal().Err(err).Msg("couldn't disconnect to mongo database")
	// 	}
	// })

	return s
}
