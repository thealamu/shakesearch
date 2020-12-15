package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type server struct {
	searcher *Searcher
	srv      *http.Server
}

func newServer(s *Searcher) *server {
	return &server{
		s,
		&http.Server{},
	}
}

func (s *server) start(addr string) error {
	fmt.Printf("Listening on %s...", addr)
	s.srv.Handler = s.routes()
	s.srv.Addr = addr

	return s.srv.ListenAndServe()
}

func (s *server) stop() error {
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return s.srv.Shutdown(shutdownCtx)
}

func (s *server) routes() http.Handler {
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/", fs)
	mux.HandleFunc("/search", handleSearch(s.searcher))

	return mux
}
