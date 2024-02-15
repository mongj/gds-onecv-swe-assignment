package router

import (
	"github.com/go-chi/chi/v5"
)

// Setup mounts all handlers
func Setup(r chi.Router) {
	// Public Routes
	r.Group(PublicRoutes())

	// A separate router for protected routes can be added here
}

func PublicRoutes() func(r chi.Router) {
	return func(r chi.Router) {
		// TODO: Add public routes here
	}
}
