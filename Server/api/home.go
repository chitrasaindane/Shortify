package api

import (
	"fmt"
	"net/http"
)

// # Serve 'Home' Page: ["/"]
// # Note: Serve the 'home' page to the 'client'
func ServeHomePage(w http.ResponseWriter, r *http.Request) {
	const rootFolderName string = "public"
	const homePageFileSubPath string = "html/index.html"

	homePageFilePath := fmt.Sprintf("%s/%s", rootFolderName, homePageFileSubPath)

	http.ServeFile(w, r, homePageFilePath)
}
