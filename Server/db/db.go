package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"shortify/env"
	"shortify/models"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// # MongoDB Client and Collection Variables
var client *mongo.Client
var collection *mongo.Collection

// # Initialize MongoDB Connection
func init() {

	// # Load Env Variables
	env.LoadEnv()

	// # Get DB User and DB Password
	dbUser := env.GetEnv("DB_USER", "User")
	dbPass := env.GetEnv("DB_PASS", "Password")

	// # Build MongoDB URI with DB User and DB Password
	mongoURI := os.Getenv("MONGO_URI")
	mongoURI = strings.Replace(mongoURI, "{DB_USER}", dbUser, 1)
	mongoURI = strings.Replace(mongoURI, "{DB_PASS}", dbPass, 1)

	// # Get DB Name and Collection Name
	dbName := env.GetEnv("DB_NAME", "Shortify")
	collectionName := env.GetEnv("COLLECTION_NAME", "Links")

	// # Connect to MongoDB
	var err error
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		// # MongoDB Connection Error
		log.Fatal("🚫 MongoDB Connection Error: ", err)
	}

	// # Ping the MongoDB server
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		// # MongoDB Ping Error
		log.Fatal("🚫 MongoDB Ping Error: ", err)
	}

	// # Set the MongoDB collection
	collection = client.Database(dbName).Collection(collectionName)

	// # MongoDB Connection Successful
	fmt.Printf("🕸️  %s Server Connected!\n", dbName)
	fmt.Printf("🕸️  %s Database Connected!\n", dbName)
}

// # SaveURL : Save a new URL to the MongoDB collection
func SaveURL(url models.URL) error {
	// # Insert URL into the MongoDB collection
	_, err := collection.InsertOne(context.TODO(), url)
	return err
}

// # GetURL : Get a URL from the MongoDB collection
func GetURL(id string) (models.URL, error) {
	var url models.URL
	// # Find URL by ID in the MongoDB collection and Decode it into URL Model Object
	err := collection.FindOne(context.TODO(), map[string]string{"id": id}).Decode(&url)
	if err != nil {
		// # URL Not Found and Returned Error
		return models.URL{}, err
	}
	// # URL Found and Returned URL
	return url, nil
}
