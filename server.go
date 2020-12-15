package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

//Server wraps a http server, providing helper methods
type Server struct {
	*http.Server
}

//Start runs the server on the provided address
func (s *Server) Start() error {
	fmt.Printf("Listening on %s...\n", s.Addr)
	return s.ListenAndServe()
}

//Stop attempts to gracefully shutdown the server
func (s *Server) Stop() error {
	fmt.Println("Shutting down")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	return s.Shutdown(shutdownCtx)
}
