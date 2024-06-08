package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/nayan9229/go-backend-services/chassis"
)

func (s *Server) routes() chi.Router {
	r := chi.NewRouter()

	// Add common middleware.
	chassis.AddCommonMiddleware(r, true)
	r.Get("/", chassis.Health)
	r.Get("/healthz", chassis.Health)
	r.Get("/html", chassis.HtmlHandler(s.HtmlHandler))
	r.Get("/json", chassis.SimpleHandler(s.JsonHandler))

	return r
}
