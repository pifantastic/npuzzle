package main

import (
	log "github.com/Sirupsen/logrus"
	"net/http"
)

func LoggingMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		log.Printf("%s: %s", request.Method, request.URL.Path)
		handler.ServeHTTP(response, request)
	})
}
