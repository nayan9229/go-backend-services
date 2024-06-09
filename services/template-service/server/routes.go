package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/nayan9229/go-backend-services/chassis"
)

func (s *Server) routes() chi.Router {
	r := chi.NewRouter()

	// Add common middleware.
	// chassis.AddCommonMiddleware(r, true)
	r.Get("/", chassis.Health)
	r.Get("/healthz", chassis.Health)
	r.Get("/html", chassis.HtmlHandler(s.HtmlHandler))
	r.Get("/json", chassis.SimpleHandler(s.JsonHandler))

	r.Route("/users", func(r chi.Router) {
		r.Get("/", chassis.SimpleHandler(s.GetUsers))
		r.Post("/", chassis.SimpleHandler(s.CreateUsers))
		r.Get("/{userID}", chassis.SimpleHandler(s.GetUserByID))
		r.Patch("/{userID}", chassis.SimpleHandler(s.UpdateUserByID))
		r.Delete("/{userID}", chassis.SimpleHandler(s.DeleteUserByID))
	})

	r.Route("/bson/users", func(r chi.Router) {
		// r.Get("/", chassis.SimpleHandler(s.GetUsers))
		r.Post("/", chassis.SimpleHandler(s.InsertUser))
		r.Get("/{userID}", chassis.SimpleHandler(s.FindUserByID))
		r.Patch("/{userID}", chassis.SimpleHandler(s.UpdateUser))
		r.Delete("/{userID}", chassis.SimpleHandler(s.DeleteUser))
	})

	return r
}
