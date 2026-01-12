package utils

import (
	"encoding/json"
	"net/http"
	"shortify/schema"
)

// # Set 'App Response'
func SetAppResponse(w http.ResponseWriter, response any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // # "200" #

	encoder := json.NewEncoder(w)

	res := &schema.AppResponse{
		Success: true,
		Payload: response,
	}

	encoder.Encode(res)
}

// # Set 'App Error'
func SetAppError(w http.ResponseWriter, errRes *schema.Error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(errRes.StatusCode) // # "error status code" #

	encoder := json.NewEncoder(w)

	err := &schema.AppError{
		Success:    false,
		Error:      errRes.Message,
		StatusCode: errRes.StatusCode,
	}

	encoder.Encode(err)
}
