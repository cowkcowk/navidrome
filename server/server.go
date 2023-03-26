package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	router *chi.Mux
	appRoot string
}

func New() *Server {
	s := &Server{}
	return s
}

