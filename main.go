package main

import (
	"fmt"
	"net/http"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	// Set content type to HTML
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// Write HTML response
	fmt.Fprintf(w, "<html><body><h1>Hello, World!</h1></body></html>")
}

func main() {
	// Handle root URL with helloWorldHandler
	http.HandleFunc("/", helloWorldHandler)

	// Start the web server on port 8080
	fmt.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
