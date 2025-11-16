package server

import (
	"log"
	"net/http"
	"time"

	"github.com/JabkaColl/go-task-6.2-sprint-final/internal/handlers"
)

type Server struct {
	HTTP   *http.Server
	logger *log.Logger
	router *http.ServeMux
}

func NewServer(logger *log.Logger) *Server {
	mux := http.NewServeMux()

	srv := &Server{
		HTTP: &http.Server{
			Addr:         ":8080",
			Handler:      mux,
			ErrorLog:     logger,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  15 * time.Second,
		},
		logger: logger,
		router: mux,
	}

	srv.registerHandlers()

	return srv
}

func (s *Server) registerHandlers() {
	s.router.HandleFunc("/", handlers.HandlerIndex)
	s.router.HandleFunc("/upload", handlers.HandlerUpload)
}
