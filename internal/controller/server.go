package controller

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Server struct {
	Server http.Server
}

func (s *Server) RunServer(port string, router *mux.Router) error {
	s.Server = http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	return s.Server.ListenAndServe()
}
