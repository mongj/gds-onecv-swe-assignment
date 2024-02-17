package router

import (
	"github.com/go-chi/chi/v5"
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
		// TODO: Add public routes here

		// Seeding
		r.Post("/seed", seed.Handler)
	}
}