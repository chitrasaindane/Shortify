package handlers

import (
	"encoding/json"
	"net/http"
	"shortify/utils"
)

// # ShortURLHandler : Create a new short URL from the original URL
func ShortURLHandler(w http.ResponseWriter, r *http.Request) {
	var data struct {
		URL string `json:"url"`
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		// # Invalid Request Body
		http.Error(w, "🚫 Invalid Request Body!", http.StatusBadRequest)
		return
	}

	// # Create a new short URL
	shortURL := utils.CreateURL(data.URL)
	response := struct {
		// # Short URL Response Object
		ShortURL string `json:"short_url"`
	}{ShortURL: shortURL}

	// # Send the response with the short URL to the client
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
