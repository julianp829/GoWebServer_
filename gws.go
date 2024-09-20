// gws.go
package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/julianp829/GoWebServer_/hello"
	"github.com/julianp829/GoWebServer_/help"
	"github.com/julianp829/GoWebServer_/syllabus"
)

func main() {
	http.HandleFunc("/hello-world", hello.HelloWorld)
	http.HandleFunc("/hello-world-html", hello.HelloWorldHTML)
	http.HandleFunc("/hello-world-json", hello.HelloWorldJSON)
	http.HandleFunc("/syllabus", syllabus.SyllabusHandler)
	http.HandleFunc("/help", help.Help)

	// Start server
	fmt.Println("Starting server on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
}
