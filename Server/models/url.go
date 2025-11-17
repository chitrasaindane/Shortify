package models

// # URL : URL Model
type URL struct {
	ID           string    `bson:"id" json:"id"`
	OriginalURL  string    `bson:"original_url" json:"original_url"`
	ShortURL     string    `bson:"short_url" json:"short_url"`
}
