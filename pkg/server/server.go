package server

import "github.com/go-chi/chi/v5"

type Server struct {
	Router *chi.Mux
	// additional server config can be added here
}

func New() *Server {
	s := &Server{}
	s.Router = chi.NewRouter()
	return s
}
