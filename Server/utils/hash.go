package utils

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"os"
	"shortify/db"
	"shortify/env"
	"shortify/models"
	"strings"
)

// # GenerateShortURL : Generate a short URL from the original URL
func GenerateShortURL(originalURL string) string {
	hasher := md5.New()
	hasher.Write([]byte(originalURL))
	hash := hex.EncodeToString(hasher.Sum(nil))
	return hash[:6]
}

// # CreateURL : Create a new URL in the database
func CreateURL(originalURL string) string {

	// # Get Short URL
	shortURL := GenerateShortURL(originalURL)
	id := shortURL

	// # Load Env Variables
	env.LoadEnv()

	// # Get Port
	port := env.GetEnv("PORT", "3000")

	// # Build Short URL with Short URL ID
	shortURL = os.Getenv("SERVER_URI")
	shortURL = strings.Replace(shortURL, "{PORT}", port, 1)
	shortURL = shortURL + "/redirect/" + id

	// # Form URL Model Object
	url := models.URL{
		ID:          id,
		OriginalURL: originalURL,
		ShortURL:    shortURL,
	}
	// # Save URL
	err := db.SaveURL(url)
	if err != nil {
		// # Saving URL Error
		log.Fatal("🚫 Saving URL Error: ", err)
	}
	// # Return Short URL
	return shortURL
}
