package api

import (
	"errors"
	"fmt"
	"net/http"
	"shortify/schema"
	"shortify/utils"
)

// # Serve 'File': ["/file"]
// # Note: Serve the requested 'file' to the 'client'
func ServeFile(w http.ResponseWriter, r *http.Request) {
	var err error
	var errMsg string
	var errRes schema.Error

	// # Get the file 'sub-path' from the 'query params'
	fileSubPath := r.URL.Query().Get("sub_path")
	fileSubPath = utils.GetTrimmedValue(fileSubPath)
	if fileSubPath == "" {
		errMsg = "🚫 Invalid file requested to shortify: Empty file sub-path provided"
		err = errors.New(errMsg)
		utils.LogError(err, "API.ServeFile")
		errRes = schema.Error{
			StatusCode: http.StatusBadRequest,
			Message:    errMsg,
		}
		utils.SetAppError(w, &errRes)
		return
	}

	// # Print the file 'sub-path'
	fmt.Println("🔗 File Sub-Path:", fileSubPath)

	const rootFolderName string = "public"

	filePath := fmt.Sprintf("%s/%s", rootFolderName, fileSubPath)

	http.ServeFile(w, r, filePath)
}
