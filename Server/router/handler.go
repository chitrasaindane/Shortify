package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

// # Allow 'CORS'
func AllowCORS(router *mux.Router) http.Handler {
	corsHandler := http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// # Set the 'CORS' headers
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusOK)
				return
			}

			router.ServeHTTP(w, r)
		},
	)

	return corsHandler
}
