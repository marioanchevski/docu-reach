package api

import (
	"log"
	"net/http"
	"time"

	"github.com/marioanchevski/docu-reach/cmd/api/document"
	"github.com/marioanchevski/docu-reach/cmd/api/health"
	"github.com/marioanchevski/docu-reach/config"
	store "github.com/marioanchevski/docu-reach/repository/document"
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

	v1 := http.NewServeMux()
	v1.Handle("/api/v1/", http.StripPrefix("/api/v1", mux))

	documentStore := store.NewInMemoryDocumentStore()
	documentHandler := document.NewHandler(documentStore)
	documentHandler.RegisterRoutes(mux)

	health := health.NewHealthHandler()
	health.RegisterRoutes(v1)

	server := &http.Server{
		Addr:         s.config.ListenAddr,
		Handler:      v1,
		WriteTimeout: time.Second * 10,
		ReadTimeout:  time.Second * 5,
		IdleTimeout:  time.Minute,
	}

	log.Printf("Starting server on port %v", s.config.ListenAddr)

	return server.ListenAndServe()
}
