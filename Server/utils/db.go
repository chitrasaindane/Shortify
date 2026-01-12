package utils

import "go.mongodb.org/mongo-driver/bson/primitive"

// # Get New 'Object ID'
func GetNewObjectID() *primitive.ObjectID {
	newObjectID := primitive.NewObjectID()
	return &newObjectID
}
