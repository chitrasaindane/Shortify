package main

import (
	"fmt"
	"log"
	"net/http"
	"shortify/env"
	"shortify/router"
	"strings"
)

// # Server "Entry" Point #
func main() {
	var err error

	// # Load the 'Env' variables
	env.LoadEnv()

	// # Get the 'Server' variables from the 'Env'
	serverURI := env.GetEnv("SERVER_URI")
	port := env.GetEnv("PORT")

	// # Build the server 'URI' with the 'port'
	serverURI = strings.Replace(serverURI, "{PORT}", port, 1)

	// # Setup the 'Router'
	r := router.SetupRouter()

	// # Allow 'CORS'
	corsHandler := router.AllowCORS(r)

	// # Start the 'server'
	fmt.Println("🚀 Server is running...")
	fmt.Println("🔗 Server Link:", serverURI)

	listenAddress := fmt.Sprintf(":%s", port)

	err = http.ListenAndServe(listenAddress, corsHandler)
	if err != nil {
		// # Start 'Server' Error
		log.Fatal("🚫 Start Server Error:", err)
	}
}
