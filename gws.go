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

func printHelp() {
	helpMessage := `
Go Web Server (GWS) Help:

This application provides the following APIs:

1. /hello-world         - Responds with plain text "Hello World – GWS"
2. /hello-world-html    - Responds with HTML "Hello World &mdash; GWS"
3. /hello-world-json    - Responds with JSON {"message": "Hello World - GWS"}
4. /syllabus            - Handles syllabus-related operations:
   - GET    /syllabus   - Returns the syllabus JSON file.
   - DELETE /syllabus   - Returns "deleted – stubbed".
   - POST   /syllabus   - Returns "create-stubbed".
   - PUT    /syllabus   - Returns "update-stubbed".
5. /d6					- Returns a random integer between 1 and 6.

Run the server by typing: go run gws.go

Example: http://localhost:8080/hello-world
`
	fmt.Println(helpMessage)
}

func main() {
	// Check if the first argument is "help"
	if len(os.Args) > 1 && os.Args[1] == "help" {
		printHelp()
		return
	}

	http.HandleFunc("/hello-world", hello.HelloWorld)
	http.HandleFunc("/hello-world-html", hello.HelloWorldHTML)
	http.HandleFunc("/hello-world-json", hello.HelloWorldJSON)
	http.HandleFunc("/syllabus", syllabus.SyllabusHandler)
	http.HandleFunc("/help", help.Help)
	http.HandleFunc("/d6", d6Handler) // Register the /d6 handler

	// Start server
	fmt.Println("Starting server on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
}
