package help

import (
	"fmt"
	"net/http"
)

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
