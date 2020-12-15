package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type Server struct {
	*http.Server
}

func (s *Server) Start() error {
	fmt.Printf("Listening on %s...", s.Addr)
	return s.ListenAndServe()
}

func (s *Server) Stop() error {
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	return s.Shutdown(shutdownCtx)
}
