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

	// API methods.
	http.HandleFunc("/api/v1/puzzles", Puzzles)
	http.HandleFunc("/api/v1/solutions", Solutions)
	http.HandleFunc("/api/v1/leaderboards", Leaderboards)

	// Assume anything else is a static file.
	http.Handle("/", http.FileServer(assetFS()))

	// Hack to allow the leaderboard page to use a different route. Renders
	// the same as the index page, but allows the UI to execute a different route.
	http.HandleFunc("/leaderboard", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	})

	// Start the web server.
	log.Printf("Listening at http://localhost:%d", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), LoggingMiddleware(http.DefaultServeMux))
}
