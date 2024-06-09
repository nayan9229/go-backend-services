package chassis

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/hlog"
	"github.com/rs/zerolog/log"
)

func AddCommonMiddleware(r chi.Router, isDev bool) {
	r.Use(hlog.NewHandler(log.Logger))

	if isDev {
		logs := func(r *http.Request, status, size int, duration time.Duration) {
			basicRequestLog(r, status, size, duration).Msg("")
		}
		r.Use(hlog.RequestIDHandler("req_id", "Request-Id"))
		r.Use(hlog.AccessHandler(logs))
		r.Use(hlog.RemoteAddrHandler("ip"))
		r.Use(hlog.UserAgentHandler("user_agent"))
		r.Use(hlog.RefererHandler("referer"))

		r.Use(logHandler)
	}
}

// Basic HTTP request logging.
func basicRequestLog(r *http.Request, status, size int, duration time.Duration) *zerolog.Event {
	if r.URL.Path == "/healthz" {
		return nil
	}
	return hlog.FromRequest(r).Info().
		Str("method", r.Method).
		Str("url", r.URL.String()).
		Int("status", status).
		Int("size", size).
		Dur("duration", duration)
}

func logHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/healthz" {
			next.ServeHTTP(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}