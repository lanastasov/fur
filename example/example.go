package main

import (
	"github.com/squiidz/fur"
	"log"
	"net/http"
)

func main() {
	// Create a NewServer
	s := fur.NewServer("localhost", ":8080", false)

	// Set some default middleware for all the handlers
	s.Stack(MiddleRedirect)

	// Add a new routes and add some middleware for this one only
	s.AddRoute("/nuts", DefaultHandler, MiddleLog)
	s.AddRoute("/", nil)

	// Start the server
	s.Start()
}

// Application Handler
func DefaultHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("Welcome to my website"))
}

// Middleware Logger
func MiddleLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		log.Printf("[%s] %s %s", req.Method, req.RequestURI, req.RemoteAddr)
		next.ServeHTTP(rw, req)
	})
}

// Useless Middleware, just for example
func MiddleRedirect(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.RequestURI != "/nuts" {
			http.Redirect(rw, req, "http://google.com", http.StatusFound)
		} else {
			next.ServeHTTP(rw, req)
		}
	})
}
