package router

import (
	"shortify/api"

	"github.com/gorilla/mux"
)

// # Setup 'Router'
func SetupRouter() *mux.Router {
	// # Get the new 'mux' router 'instance'
	router := mux.NewRouter()

	// # Serve 'Home' Page: ["/"]
	// # Note: Serve the 'home' page to the 'client'
	router.HandleFunc("/", api.ServeHomePage).Methods("GET")

	// # Serve 'File': ["/file"]
	// # Note: Serve the requested 'file' to the 'client'
	router.HandleFunc("/file", api.ServeFile).Methods("GET")

	// # Redirect 'Route'
	// # Note: Redirect the 'user' to the 'Destination URL' of the 'link'
	router.HandleFunc("/{username}/{slug}", api.RedirectURL).Methods("GET")

	// # API 'Base Path'
	APIBasePath := "/api/v1"

	// # 'API' Version '1' Router
	APIV1 := router.PathPrefix(APIBasePath).Subrouter()

	// # 'API' Version '1' Sub 'Routers'
	UserV1 := APIV1.PathPrefix("/user").Subrouter()
	LinkV1 := APIV1.PathPrefix("/link").Subrouter()

	// # 'User' Routes
	UserV1.HandleFunc("/webhook", api.HandleClerkUserWebhook).Methods("POST")
	UserV1.HandleFunc("/username", api.UpdateUsername).Methods("PUT")
	UserV1.HandleFunc("/", api.GetUser).Methods("GET")
	UserV1.HandleFunc("/list", api.GetUsers).Methods("GET")

	// # 'Link' Routes
	LinkV1.HandleFunc("/", api.CreateLink).Methods("POST")
	LinkV1.HandleFunc("/", api.UpdateLink).Methods("PUT")
	LinkV1.HandleFunc("/", api.DeleteLink).Methods("DELETE")
	LinkV1.HandleFunc("/", api.GetLink).Methods("GET")
	LinkV1.HandleFunc("/list", api.GetLinks).Methods("GET")

	// # Return the 'configured' mux 'router' instance
	return router
}
