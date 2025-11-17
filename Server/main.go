package main

import (
	"fmt"
	"net/http"
	"os"
	"shortify/env"
	"shortify/handlers"
	"strings"
)

func main() {

	// # Load Env Variables
	env.LoadEnv()

	// # Get Port
	port := env.GetEnv("PORT", "3000")

	// # Build Server URL with Port
	serverURI := os.Getenv("SERVER_URI")
	serverURI = strings.Replace(serverURI, "{PORT}", port, 1)

	// # Routes
	http.HandleFunc("/", handlers.RootPageURL)
	http.HandleFunc("/shorten", handlers.ShortURLHandler)
	http.HandleFunc("/redirect/", handlers.RedirectURLHandler)

	// # Start Server
	fmt.Printf("🚀 Server is running...\n")
	fmt.Printf("🔗 Link : %s\n", serverURI)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		// # Start Server Error
		fmt.Printf("🚫 Start Server Error : %v\n", err)
	}
}
