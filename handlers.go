// handlers.go
package main

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

// Help provides API instructions
func Help(w http.ResponseWriter, r *http.Request) {
	helpMessage := `
	<h1>API Help</h1>
	<ul>
		<li><a href="/hello-world">/hello-world</a></li>
		<li><a href="/hello-world-html">/hello-world-html</a></li>
		<li><a href="/hello-world-json">/hello-world-json</a></li>
		<li><a href="/syllabus">/syllabus</a></li>
	</ul>`
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintln(w, helpMessage)
}
