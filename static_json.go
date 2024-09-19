package main

import (
	_ "embed"
	"net/http"
)

// Embed the JSON file using Go's "embed" package
//
//go:embed syllabus.json
var syllabusData []byte

func helloWorldJSONStaticHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(syllabusData)
}
