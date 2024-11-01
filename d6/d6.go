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
