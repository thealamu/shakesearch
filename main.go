package main

import (
	"log"
	"net"
	"os"
)

func main() {
	searcher := &Searcher{}
	err := searcher.Load("completeworks.txt")
	if err != nil {
		log.Fatal(err)
	}

	srv := newServer(searcher)
	err = srv.start(getRunAddr())
	if err != nil {
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
