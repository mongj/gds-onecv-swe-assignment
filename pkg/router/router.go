package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/mongj/gds-onecv-swe-assignment/pkg/handlers"
	"github.com/mongj/gds-onecv-swe-assignment/pkg/seed"
)

// Setup mounts all subrouters and handlers
func Setup(r chi.Router) {
	// Subrouter for /api path
	apiRouter := chi.NewRouter()
	apiRouter.Group(PublicRoutes())
	r.Mount("/api", apiRouter)
}

func PublicRoutes() func(r chi.Router) {
	return func(r chi.Router) {
		// Handlers for each user story
		r.Post("/register", handlers.RegisterStudent)
		r.Get("/commonstudents", handlers.ListCommonStudents)
		r.Post("/suspend", handlers.SuspendStudent)
		r.Post("/retrievefornotifications", handlers.RetrieveForNotifications)

		// Seeding
		r.Post("/seed", seed.Handler)
	}
}
