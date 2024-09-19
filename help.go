package main

import (
	"fmt"
	"net/http"
)

func helpHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, `
		<h1>Go Web Server API Help</h1>
		<ul>
			<li><a href="/hello-world">/hello-world</a> - Plain text Hello World</li>
			<li><a href="/hello-world-html">/hello-world-html</a> - HTML Hello World</li>
			<li><a href="/hello-world-json">/hello-world-json</a> - JSON Hello World</li>
			<li><a href="/hello-world-json-static">/hello-world-json-static</a> - Embedded JSON response</li>
			<li><a href="/syllabus">/syllabus</a> - Syllabus API (CRUD operations)</li>
		</ul>
	`)
}
