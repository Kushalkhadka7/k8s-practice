package server

import (
	"auth/config"
	"auth/router"
	"fmt"
	"log"
	"net/http"
)

// Server holds configurations for a http server.
type Server struct {
	*http.Server
}

// Creator initializes server interface.
type Creator interface {
	Start() error
}

// New initializes new http server.
func New(c *config.Config, router *router.Router) Creator {

	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", c.Server.Port),
		Handler:      router,
		ReadTimeout:  c.Server.ReadTimeOut,
		WriteTimeout: c.Server.WriteTimeOut,
	}

	return &Server{
		s,
	}
}

// Start http server in given port.
func (s *Server) Start() error {
	log.Printf("Starting server on port %s...", s.Addr)

	err := s.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
