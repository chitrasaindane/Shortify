package app

import (
	"context"
	"errors"
	"shortify/db"
	"shortify/model"
	"shortify/schema"
	"shortify/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// # Create 'Link'
func CreateLink(clerkUserID string, link *schema.Link) (*string, error) {
	ctx := context.TODO()

	var err error
	var errMsg string
	var successMsg string

	// # Base 'Filter'
	filter := bson.M{
		"clerk_user_id": clerkUserID,
		"slug":          utils.GetStringValue(link.Slug),
	}

	// # Check if the provided 'slug' is already 'assigned' to another 'link' of the given 'user'
	Link, err := GetLink(filter)
	if Link != nil {
		errMsg = "The provided slug is already assigned to another link. Please choose a different slug."
		err = errors.New(errMsg)
		return nil, err
	}

	// # Base 'Filter'
	filter = bson.M{
		"clerk_user_id": clerkUserID,
	}

	// # Get the 'user'
	user, err := GetUser(filter)
	if err != nil {
		utils.LogError(err, "App.CreateLink")
		return nil, err
	}

	collection := db.GetMongoCollection(model.LinkColl)

	// # Get the 'current system time' in the 'IST' timezone
	time, err := utils.GetCurrentSystemTimeInIST()
	if err != nil {
		utils.LogError(err, "App.CreateLink")
		return nil, err
	}

	Link = &model.Link{
		ID:             utils.GetNewObjectID(),
		UserID:         user.ID,
		ClerkUserID:    user.ClerkUserID,
		IsUserDeleted:  false,
		Title:          utils.GetStringValue(link.Title),
		Description:    utils.GetStringValue(link.Description),
		DestinationURL: utils.GetStringValue(link.DestinationURL),
		Slug:           utils.GetStringValue(link.Slug),
		CreatedAt:      time,
		CreatedBy:      utils.GetActionUser(clerkUserID),
	}

	// # Insert the 'new link'
	_, err = collection.InsertOne(ctx, Link)
	if err != nil {
		utils.LogError(err, "App.CreateLink")
		errMsg = "Failed to create the new link"
		err = errors.New(errMsg)
		return nil, err
	}

	successMsg = "Link created successfully!"

	return &successMsg, nil
}

// # Update 'Link'
func UpdateLink(clerkUserID string, linkID *primitive.ObjectID, link *schema.Link) (*string, error) {
	ctx := context.TODO()

	var err error
	var errMsg string
	var successMsg string

	// # Base 'Filter'
	filter := bson.M{
		"clerk_user_id": clerkUserID,
		"slug":          utils.GetStringValue(link.Slug),
	}

	// # Check if the provided 'slug' is already 'assigned' to another 'link' of the given 'user'
	Link, err := GetLink(filter)
	if Link != nil && Link.ID.Hex() != linkID.Hex() {
		errMsg = "The provided slug is already assigned to another link. Please choose a different slug."
		err = errors.New(errMsg)
		return nil, err
	}

	// # Base 'Filter'
	filter = bson.M{
		"_id": linkID,
	}

	collection := db.GetMongoCollection(model.LinkColl)

	// # Get the 'current system time' in the 'IST' timezone
	time, err := utils.GetCurrentSystemTimeInIST()
	if err != nil {
		utils.LogError(err, "App.UpdateLink")
		return nil, err
	}

	// # Base 'Set' & 'Unset'
	set, unset := bson.M{}, bson.M{}

	if link.Title != nil {
		if utils.GetStringValue(link.Title) != "" {
			set["title"] = utils.GetStringValue(link.Title)
		} else {
			unset["title"] = 1
		}
	}
	if link.Description != nil {
		if utils.GetStringValue(link.Description) != "" {
			set["description"] = utils.GetStringValue(link.Description)
		} else {
			unset["description"] = 1
		}
	}
	if link.DestinationURL != nil {
		if utils.GetStringValue(link.DestinationURL) != "" {
			set["destination_url"] = utils.GetStringValue(link.DestinationURL)
		} else {
			unset["destination_url"] = 1
		}
	}
	if link.Slug != nil {
		if utils.GetStringValue(link.Slug) != "" {
			set["slug"] = utils.GetStringValue(link.Slug)
		} else {
			unset["slug"] = 1
		}
	}

	setLength := len(set)
	unsetLength := len(unset)

	if setLength > 0 || unsetLength > 0 {
		set["updated_at"] = time
		set["updated_by"] = utils.GetActionUser(clerkUserID)

		// # Base 'Update'
		update := bson.M{
			"$set": set,
		}

		if unsetLength > 0 {
			update["$unset"] = unset
		}

		// # Update the 'link'
		res, err := collection.UpdateOne(ctx, filter, update)
		if err != nil {
			utils.LogError(err, "App.UpdateLink")
			errMsg = "Failed to update the link"
			err = errors.New(errMsg)
			return nil, err
		}
		if res.MatchedCount == 0 {
			errMsg = "Link not found"
			err = errors.New(errMsg)
			return nil, err
		}
		if res.ModifiedCount == 0 {
			errMsg = "No link updated"
			err = errors.New(errMsg)
			return nil, err
		}
	}

	successMsg = "Link updated successfully!"

	return &successMsg, nil
}

// # Delete 'Link'
func DeleteLink(linkID *primitive.ObjectID) (*string, error) {
	ctx := context.TODO()

	var err error
	var errMsg string
	var successMsg string

	collection := db.GetMongoCollection(model.LinkColl)

	// # Base 'Filter'
	filter := bson.M{
		"_id": linkID,
	}

	// # Delete the 'link'
	res, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		utils.LogError(err, "App.DeleteLink")
		errMsg = "Failed to delete the link"
		err = errors.New(errMsg)
		return nil, err
	}
	if res.DeletedCount == 0 {
		errMsg = "Link not found"
		err = errors.New(errMsg)
		return nil, err
	}

	successMsg = "Link deleted successfully!"

	return &successMsg, nil
}

// # Get 'Link'
func GetLink(filter bson.M) (*model.Link, error) {
	ctx := context.TODO()

	var err error
	var errMsg string

	collection := db.GetMongoCollection(model.LinkColl)

	var link model.Link

	// # Get the 'link'
	err = collection.FindOne(ctx, filter).Decode(&link)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			utils.LogError(err, "App.GetLink")
			errMsg = "Link not found"
			err = errors.New(errMsg)
			return nil, err
		} else {
			utils.LogError(err, "App.GetLink")
			errMsg = "Failed to get the link"
			err = errors.New(errMsg)
			return nil, err
		}
	}
	if link.IsUserDeleted {
		errMsg = "User associated with the given link is already deleted"
		err = errors.New(errMsg)
		return nil, err
	}

	return &link, nil
}

// # Get 'Links'
func GetLinks(clerkUserID string) ([]*model.Link, error) {
	ctx := context.TODO()

	var err error
	var errMsg string

	collection := db.GetMongoCollection(model.LinkColl)

	// # Base 'Filter'
	filter := bson.M{
		"clerk_user_id": clerkUserID,
		"is_user_deleted": bson.M{
			"$ne": true, // # Exclude the 'deleted' users
		},
	}

	// # Base 'Sort'
	sort := bson.M{
		"created_at": -1, // # Sort the 'documents' (links) by the 'created_at' field in the 'descending' order
	}

	// # Set the 'Find' options
	opts := options.Find()
	opts.SetSort(sort)

	var links []*model.Link

	// # Get the 'links'
	cur, err := collection.Find(ctx, filter, opts)
	if err != nil {
		utils.LogError(err, "App.GetLinks")
		errMsg = "Failed to get the links"
		err = errors.New(errMsg)
		return nil, err
	}

	err = cur.All(ctx, &links)
	if err != nil {
		utils.LogError(err, "App.GetLinks")
		errMsg = "Failed to decode the links"
		err = errors.New(errMsg)
		return nil, err
	}

	return links, nil
}
