package handlers

import (
	"fmt"
	"net/http"
)

// # RootPageURL handles requests to the root URL
func RootPageURL(w http.ResponseWriter, r *http.Request) {
	// # Send a welcome message to the client
	fmt.Fprintf(w, "🚀 Welcome to Shortify 🚀")
}
