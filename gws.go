// gws.go
package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/hello-world", HelloWorld)
	http.HandleFunc("/hello-world-html", HelloWorldHTML)
	http.HandleFunc("/hello-world-json", HelloWorldJSON)
	http.HandleFunc("/syllabus", SyllabusHandler)
	http.HandleFunc("/help", Help)

	// Start server
	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
}
