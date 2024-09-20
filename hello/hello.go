package hello

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// HelloWorld responds with plain text
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World – GWS")
}

// HelloWorldHTML responds with HTML
func HelloWorldHTML(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintln(w, "<h1>Hello World &mdash; GWS</h1>")
}

// HelloWorldJSON responds with a JSON message
func HelloWorldJSON(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "Hello World - GWS"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
