// syllabus.go
package syllabus

import (
	"embed"
	"fmt"
	"net/http"
)

//go:embed syllabus.json
var syllabusFile embed.FS

// SyllabusHandler handles the /syllabus API
func SyllabusHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// Read embedded JSON file and return it
		data, err := syllabusFile.ReadFile("syllabus.json")
		if err != nil {
			http.Error(w, "Error reading syllabus", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	case "DELETE":
		// Stubbed DELETE implementation
		fmt.Fprintln(w, "deleted – stubbed")
	case "POST":
		// Stubbed POST (Create) implementation
		fmt.Fprintln(w, "create-stubbed")
	case "PUT":
		// Stubbed PUT (Update) implementation
		fmt.Fprintln(w, "update-stubbed")
	default:
		http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
	}
}
