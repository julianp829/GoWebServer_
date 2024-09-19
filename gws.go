package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// Routes
	http.HandleFunc("/hello-world", helloWorldHandler)
	http.HandleFunc("/hello-world-html", helloWorldHTMLHandler)
	http.HandleFunc("/hello-world-json", helloWorldJSONHandler)
	http.HandleFunc("/hello-world-json-static", helloWorldJSONStaticHandler)
	http.HandleFunc("/syllabus", syllabusHandler)
	http.HandleFunc("/help", helpHandler)

	// Handle the "gws help" command line argument
	if len(os.Args) > 1 && os.Args[1] == "help" {
		printHelp()
		return
	}

	// Start the server
	fmt.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

// Prints help instructions for the command-line version
func printHelp() {
	fmt.Println("Usage: go run gws.go")
	fmt.Println("API Endpoints:")
	fmt.Println(" - /hello-world")
	fmt.Println(" - /hello-world-html")
	fmt.Println(" - /hello-world-json")
	fmt.Println(" - /hello-world-json-static")
	fmt.Println(" - /syllabus [GET, POST, PUT, DELETE]")
}
