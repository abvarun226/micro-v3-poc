package handler

import (
	"net/http"
)

// HelloWorld handler.
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	// Make a call to grpc service here.
	// Get data from the grpc service.
	// Use api model to convert to json and render the json response.
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(`{"hello":"world"}`))
}
