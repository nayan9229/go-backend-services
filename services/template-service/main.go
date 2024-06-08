package main

import (
	"github.com/joeshaw/envdecode"
	"github.com/nayan9229/go-backend-services/services/template-service/server"
	"github.com/rs/zerolog/log"
)

var appname = "template-service"

var release = "0.0.1"

func main() {
	var cfg server.Config

	err := envdecode.StrictDecode(&cfg)
	if err != nil {
		log.Fatal().Err(err).
			Msg("failed to process environment variables")
	}

	cfg.AppName = appname
	cfg.Release = release
	serv := server.NewServer(&cfg)
	serv.Serve()
}
