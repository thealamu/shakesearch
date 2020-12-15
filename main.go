package main

import (
	"errors"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/sethvargo/go-signalcontext"
)

func main() {
	ctx, cancel := signalcontext.OnInterrupt()
	defer cancel()

	searcher := &Searcher{}
	err := searcher.Load("completeworks.txt")
	if err != nil {
		log.Fatal(err)
	}

	srv := &Server{
		&http.Server{
			Handler: getRoutes(searcher),
			Addr:    getRunAddr(),
		},
	}

	go func() {
		err := srv.Start()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()

	//wait for an interrupt
	<-ctx.Done()

	//stop the server
	if err = srv.Stop(); err != nil {
		log.Fatal(err)
	}
}

func getRunAddr() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}
	return net.JoinHostPort("", port)
}

func getRoutes(searcher *Searcher) http.Handler {
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/", fs)
	mux.HandleFunc("/search", handleSearch(searcher))

	return mux
}
