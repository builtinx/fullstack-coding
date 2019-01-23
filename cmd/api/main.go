package main

import (
	"log"
	"net/http"

	"github.com/builtinx/fullstack-exercise/pkg/hello"
	"github.com/gobuffalo/pop"
	"github.com/gorilla/mux"
)

func main() {
	rtr := mux.NewRouter()

	// When you want to use the connection to the database, simply
	// rename `_` to `db` (or something else!).
	_, err := pop.Connect("development")
	if err != nil {
		log.Fatal(err)
	}

	// Some middleware for headers
	rtr.Use(jsonContentType, corsWildcard)

	// Our default route
	rtr.HandleFunc("/", hello.World)

	// And now we launch our server
	server := &http.Server{
		Addr:    "0.0.0.0:8000",
		Handler: rtr,
	}

	log.Println("Server is up and awaiting connections to port 8000")
	log.Fatal(server.ListenAndServe())
}

// jsonContentType implements a middleware function which sets the
// content-type to JSON.
func jsonContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

// corsWildcard implements middleware to set the CORS header to allow
// any incoming origin header.
func corsWildcard(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}
