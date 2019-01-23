package hello

import "net/http"

// World implements a basic "Hello world" HTTP handler.
func World(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"message":"Hello, world!"}`))
}
