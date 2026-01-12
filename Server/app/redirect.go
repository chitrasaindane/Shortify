package app

import (
	"shortify/utils"

	"go.mongodb.org/mongo-driver/bson"
)

// # Get the 'Destination URL'
func GetDestinationURL(username, linkSlug string) (*string, error) {
	var err error

	// # Base 'User Filter'
	userFilter := bson.M{
		"username": username,
	}

	// # Get the 'user'
	user, err := GetUser(userFilter)
	if err != nil {
		utils.LogError(err, "App.GetDestinationURL")
		return nil, err
	}

	linkFilter := bson.M{
		"clerk_user_id": user.ClerkUserID,
		"slug":          linkSlug,
	}

	// # Get the 'link'
	link, err := GetLink(linkFilter)
	if err != nil {
		utils.LogError(err, "App.GetDestinationURL")
		return nil, err
	}
	destinationURL := link.DestinationURL

	return &destinationURL, nil
}
