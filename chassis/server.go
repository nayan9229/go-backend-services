package chassis

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)

type Server struct {
	AppName  string
	Ctx      context.Context
	Srv      *http.Server
	muAtExit sync.Mutex
	atExit   []func()
}

type ServerConfig struct {
	AppName        string
	Project        string `env:"PROJECT_ID,default=dev"`
	Port           int    `env:"PORT,default=9002"`
	Credentials    string `env:"CREDENTIALS_PATH"`
	SentryDSN      string `env:"SENTRY_DSN"`
	SentryTracing  bool   `env:"SENTRY_TRACING,default=false"`
	Environment    string `env:"ENVIRONMENT,default=dev"`
	MaxConnections int    `env:"MAX_CONNECTIONS,default=50"`
	Release        string
}

func (s *Server) Init(cfg *ServerConfig, r chi.Router) {
	s.AppName = cfg.AppName
	s.Ctx = context.Background()
	s.atExit = []func(){}
	s.Srv = &http.Server{
		DisableGeneralOptionsHandler: false,
		Addr:                         net.JoinHostPort("", fmt.Sprintf("%d", cfg.Port)),
		Handler:                      r,
		ReadTimeout:                  30 * time.Second,
		ReadHeaderTimeout:            10 * time.Second,
		WriteTimeout:                 30 * time.Second,
		IdleTimeout:                  30 * time.Second,
	}
}

func (s *Server) Serve() {
	errChan := make(chan error, 1)
	go func() {
		log.Info().
			Str("address", s.Srv.Addr).
			Msg("server started")
		err := s.Srv.ListenAndServe()
		if err != nil {
			errChan <- err
		}
	}()

	signalCh := make(chan os.Signal, 2)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)

	var err error

	select {
	case <-signalCh:
	case err = <-errChan:
	}

	s.shutdown()
	s.runAtShutdown()

	if err != nil {
		log.Error().Err(err).Msg("server failed")
	}
}

func (s *Server) shutdown() {
	ctx, cancel := context.WithTimeout(s.Ctx, 10*time.Second)
	defer cancel()
	if err := s.Srv.Shutdown(ctx); err != nil {
		log.Error().Err(err)
	}
}

// AddAtExit adds an exit handler function.
func (s *Server) AddAtExit(fn func()) {
	s.muAtExit.Lock()
	defer s.muAtExit.Unlock()
	s.atExit = append(s.atExit, fn)
}

// AtExit executes all registered exit handlers.
func (s *Server) runAtShutdown() {
	s.muAtExit.Lock()
	defer s.muAtExit.Unlock()
	for _, fn := range s.atExit {
		fn()
	}
}

type SimpleHandlerFunc func(w http.ResponseWriter, r *http.Request) (interface{}, error)

func SimpleHandler(inner SimpleHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result, err := inner(w, r)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Printf("handling %q: %v", r.RequestURI, err)
			return
		}

		if result == nil {
			return
		}

		// Marshal JSON response body.
		body, err := json.Marshal(result)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Printf("handling %q: %v", r.RequestURI, err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
	}
}

func HtmlHandler(inner SimpleHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result, err := inner(w, r)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Printf("handling %q: %v", r.RequestURI, err)
			return
		}

		if result == nil {
			return
		}

		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(result.(string)))
	}
}
