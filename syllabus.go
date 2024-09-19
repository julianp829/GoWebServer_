package main

import (
	"fmt"
	"net/http"
)

type Syllabus struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func syllabusHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// Return the embedded JSON syllabus
		w.Header().Set("Content-Type", "application/json")
		w.Write(syllabusData)
	case http.MethodDelete:
		// Respond with stubbed delete message
		fmt.Fprintln(w, "deleted – stubbed")
	case http.MethodPost:
		// Respond with stubbed create message
		fmt.Fprintln(w, "create-stubbed")
	case http.MethodPut:
		// Respond with stubbed update message
		fmt.Fprintln(w, "update-stubbed")
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
