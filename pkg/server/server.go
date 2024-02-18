package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/mongj/gds-onecv-swe-assignment/pkg/middleware"
	"github.com/mongj/gds-onecv-swe-assignment/pkg/router"
)

type Server struct {
	Router *chi.Mux
	// additional server config can be added here
}

func New() *Server {
	s := &Server{}
	s.Router = chi.NewRouter()
	return s
}

func (s *Server) MountHandlers() {
	router.Setup(s.Router)
}

func (s *Server) MountMiddleware() {
	middleware.Setup(s.Router)
}
