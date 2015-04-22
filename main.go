package main

import (
	"flag"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"net/http"
)

const (
	version = "0.0.1"
)

var (
	port    uint
	rootDir = "."
)

func init() {
	flag.UintVar(&port, "port", 3000, "the port to listen on")
	flag.UintVar(&port, "p", 3000, "the port to listen on")
}

func main() {
	flag.Parse()

	http.HandleFunc("/api/v1/puzzles", Puzzles)
	http.HandleFunc("/api/v1/solutions", Solutions)
	http.HandleFunc("/api/v1/leaderboards", Leaderboards)

	http.Handle("/", http.FileServer(assetFS()))

	http.HandleFunc("/leaderboard", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	})

	log.Printf("Listening at http://localhost:%d", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), LoggingMiddleware(http.DefaultServeMux))
}
