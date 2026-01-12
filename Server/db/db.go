package db

import (
	"context"
	"fmt"
	"log"
	env "shortify/env"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// # MongoDB 'Variables'
var client *mongo.Client
var database *mongo.Database
var collection *mongo.Collection

// # Initialize 'MongoDB' Connection
func init() {
	ctx := context.TODO()

	var err error

	// # Load the 'Env' variables
	env.LoadEnv()

	fmt.Println("#----------| 🚀 Welcome to Shortify - Modern URL Shortener 🚀 |----------#")

	// # Get the 'DB' variables from the 'Env'
	dbName := env.GetEnv("DB_NAME")
	dbUser := env.GetEnv("DB_USER")
	dbPass := env.GetEnv("DB_PASS")

	// # Build the 'MongoDB URI'
	mongoURI := env.GetEnv("MONGO_URI")
	mongoURI = strings.Replace(mongoURI, "{DB_USER}", dbUser, 1)
	mongoURI = strings.Replace(mongoURI, "{DB_PASS}", dbPass, 1)
	mongoURI = strings.Replace(mongoURI, "{DB_NAME}", dbName, 1)

	// # Connect to the 'MongoDB'
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		// # MongoDB 'Connection' Error
		log.Fatal("🚫 MongoDB Connection Error:", err)
	}

	// # Get the 'MongoDB' database
	database = client.Database(dbName)

	// # Ping the 'MongoDB' server
	err = client.Ping(ctx, nil)
	if err != nil {
		// # MongoDB 'Ping' Error
		log.Fatal("🚫 MongoDB Ping Error:", err)
	}

	// # MongoDB 'Connection' Successful
	fmt.Println("🕸️  Server Connected!")
	fmt.Println("🕸️  Database Connected!")
}

// # Get 'MongoDB' Collection
func GetMongoCollection(collectionName string) *mongo.Collection {
	collection = database.Collection(collectionName)
	return collection
}
