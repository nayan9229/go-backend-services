package server

import "github.com/nayan9229/go-backend-services/chassis"

type Server struct {
	chassis.Server
	stageEnv bool
}

type Config struct {
	chassis.ServerConfig
	DevMode bool `env:"DEV_MODE,default=false"`
}

func NewServer(cfg *Config) *Server {
	s := &Server{
		stageEnv: cfg.DevMode,
	}
	s.Init(&cfg.ServerConfig, s.routes())
	return s
}
