// d6.go
package d6

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

// d6Handler generates and returns a random integer between 1 and 6.
func d6Handler(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator
	randomNumber := rand.Intn(6) + 1 // Generate a number between 1 and 6
	fmt.Fprintf(w, "%d", randomNumber)
}

func main() {
	// Register the /d6 endpoint with the handler function
	http.HandleFunc("/d6", d6Handler)

	// Start the server on port 8080
	fmt.Println("Starting d6 server on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
