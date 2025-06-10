package api

import (
	"log"
	"net/http"
	"time"

	"github.com/marioanchevski/docu-reach/config"
)

type APIServer struct {
	config *config.Config
}

func NewAPIServer(cfg *config.Config) *APIServer {
	return &APIServer{
		config: cfg,
	}
}

func (s *APIServer) Run() error {
	mux := http.NewServeMux()

	server := &http.Server{
		Addr:         s.config.ListenAddr,
		Handler:      mux,
		WriteTimeout: time.Second * 10,
		ReadTimeout:  time.Second * 5,
		IdleTimeout:  time.Minute,
	}

	log.Printf("Starting server on port %v", s.config.ListenAddr)

	return server.ListenAndServe()
}
