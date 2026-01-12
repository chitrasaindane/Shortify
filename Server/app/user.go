package app

import (
	"context"
	"errors"
	"fmt"
	"shortify/db"
	"shortify/model"
	"shortify/schema"
	"shortify/utils"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// # Generate 'User' Code
func GenerateUserCode() (*int64, error) {
	ctx := context.TODO()

	var err error
	var errMsg string

	collection := db.GetMongoCollection(model.UserColl)

	// # Base 'Filter'
	filter := bson.M{} // # "Empty" Filter

	// # Base 'Sort'
	sort := bson.M{
		"code": -1, // # Sort the 'documents' (users) by the 'code' field in the 'descending' order
	}

	// # Set the 'FindOne' options
	opts := options.FindOne()
	opts.SetSort(sort)

	var user *model.User
	var userCode int64

	// # Get the 'last user'
	err = collection.FindOne(ctx, filter, opts).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			userCode = 1
		} else {
			utils.LogError(err, "App.GenerateUserCode")
			errMsg = "Failed to get the last user"
			err = errors.New(errMsg)
			return nil, err
		}
	} else {
		userCode = (user.Code + 1)
	}

	return &userCode, nil
}

// # Create 'User'
func CreateUser(clerkUser *schema.ClerkUser) (*string, error) {
	ctx := context.TODO()

	var err error
	var errMsg string
	var successMsg string

	collection := db.GetMongoCollection(model.UserColl)

	// # Get the 'current system time' in the 'IST' timezone
	time, err := utils.GetCurrentSystemTimeInIST()
	if err != nil {
		utils.LogError(err, "App.CreateUser")
		return nil, err
	}

	user := "clerk_webhook"

	User := model.User{
		ID:          utils.GetNewObjectID(),
		ClerkUserID: utils.GetTrimmedValue(clerkUser.ID),
		FirstName:   utils.GetStringValue(clerkUser.FirstName),
		LastName:    utils.GetStringValue(clerkUser.LastName),
		Email:       utils.GetStringValue(clerkUser.EmailAddresss[0].EmailAddress),
		CreatedAt:   time,
		CreatedBy:   utils.GetActionUser(user),
	}

	// # Generate the 'user code'
	userCode, err := GenerateUserCode()
	if err != nil {
		utils.LogError(err, "App.CreateUser")
		errMsg = "Failed to generate the user code"
		err = errors.New(errMsg)
		return nil, err
	}
	User.Code = *userCode

	username := fmt.Sprintf("%s%s%d", strings.ToLower(User.FirstName), strings.ToLower(User.LastName), User.Code)
	User.Username = username

	// # Insert the 'new user'
	_, err = collection.InsertOne(ctx, &User)
	if err != nil {
		utils.LogError(err, "App.CreateUser")
		errMsg = "Failed to create the new user"
		err = errors.New(errMsg)
		return nil, err
	}

	successMsg = "User created successfully!"

	return &successMsg, nil
}

// # Update 'User'
func UpdateUser(clerkUser *schema.ClerkUser) (*string, error) {
	ctx := context.TODO()

	var err error
	var errMsg string
	var successMsg string

	collection := db.GetMongoCollection(model.UserColl)

	// # Get the 'current system time' in the 'IST' timezone
	time, err := utils.GetCurrentSystemTimeInIST()
	if err != nil {
		utils.LogError(err, "App.UpdateUser")
		return nil, err
	}

	user := "clerk_webhook"

	// # Base 'Filter'
	filter := bson.M{
		"clerk_user_id": utils.GetTrimmedValue(clerkUser.ID),
	}

	// # Base 'Set' & 'Unset'
	set, unset := bson.M{}, bson.M{}

	if clerkUser.FirstName != nil {
		if utils.GetStringValue(clerkUser.FirstName) != "" {
			set["first_name"] = utils.GetStringValue(clerkUser.FirstName)
		} else {
			unset["first_name"] = 1
		}
	}
	if clerkUser.LastName != nil {
		if utils.GetStringValue(clerkUser.LastName) != "" {
			set["last_name"] = utils.GetStringValue(clerkUser.LastName)
		} else {
			unset["last_name"] = 1
		}
	}
	if clerkUser.EmailAddresss[0].EmailAddress != nil {
		if utils.GetStringValue(clerkUser.EmailAddresss[0].EmailAddress) != "" {
			set["email"] = utils.GetStringValue(clerkUser.EmailAddresss[0].EmailAddress)
		} else {
			unset["email"] = 1
		}
	}
	if clerkUser.ImageURL != nil {
		if utils.GetStringValue(clerkUser.ImageURL) != "" {
			set["image_url"] = utils.GetStringValue(clerkUser.ImageURL)
		} else {
			unset["image_url"] = 1
		}
	}

	setLength := len(set)
	unsetLength := len(unset)

	if setLength > 0 || unsetLength > 0 {
		set["updated_at"] = time
		set["updated_by"] = utils.GetActionUser(user)

		// # Base 'Update'
		update := bson.M{
			"$set": set,
		}

		if unsetLength > 0 {
			update["$unset"] = unset
		}

		// # Update the 'user'
		res, err := collection.UpdateOne(ctx, filter, update)
		if err != nil {
			utils.LogError(err, "App.UpdateUser")
			errMsg = "Failed to update the user"
			err = errors.New(errMsg)
			return nil, err
		}
		if res.MatchedCount == 0 {
			errMsg = "User not found"
			err = errors.New(errMsg)
			return nil, err
		}
		if res.ModifiedCount == 0 {
			errMsg = "No user updated"
			err = errors.New(errMsg)
			return nil, err
		}
	}

	successMsg = "User updated successfully!"

	return &successMsg, nil
}

// # Delete 'User'
func DeleteUser(clerkUser *schema.ClerkUser) (*string, error) {
	ctx := context.TODO()

	var err error
	var errMsg string
	var successMsg string

	useeCollection := db.GetMongoCollection(model.UserColl)
	linkCollection := db.GetMongoCollection(model.LinkColl)

	// # Get the 'current system time' in the 'IST' timezone
	time, err := utils.GetCurrentSystemTimeInIST()
	if err != nil {
		utils.LogError(err, "App.DeleteUser")
		return nil, err
	}

	user := "clerk_webhook"

	// # 1. Mark the given 'user' as 'deleted'

	// # Base 'Filter'
	filter := bson.M{
		"clerk_user_id": utils.GetTrimmedValue(clerkUser.ID),
	}

	// # Base 'User Set'
	userSet := bson.M{
		"is_deleted": clerkUser.Deleted,
		"deleted_at": time,
		"deleted_by": utils.GetActionUser(user),
	}

	// # Base 'User Update'
	userUpdate := bson.M{
		"$set": userSet,
	}

	// # Mark the given 'user' as 'deleted'
	res, err := useeCollection.UpdateOne(ctx, filter, userUpdate)
	if err != nil {
		utils.LogError(err, "App.DeleteUser")
		errMsg = "Failed to mark the given user as deleted"
		err = errors.New(errMsg)
		return nil, err
	}
	if res.MatchedCount == 0 {
		errMsg = "User not found"
		err = errors.New(errMsg)
		return nil, err
	}
	if res.ModifiedCount == 0 {
		errMsg = "No user marked as deleted"
		err = errors.New(errMsg)
		return nil, err
	}

	// # 2. Mark all the 'links' of the given 'user' as 'deleted'

	// # Base 'Link Set'
	linkSet := bson.M{
		"is_user_deleted": clerkUser.Deleted,
		"updated_at":      time,
		"updated_by":      utils.GetActionUser(user),
	}

	// # Base 'Link Update'
	linkUpdate := bson.M{
		"$set": linkSet,
	}

	// # Mark all the 'links' of the given 'user' as 'deleted'
	res, err = linkCollection.UpdateMany(ctx, filter, linkUpdate)
	if err != nil {
		utils.LogError(err, "App.DeleteUser")
		errMsg = "Failed to mark all the links of the given user as deleted"
		err = errors.New(errMsg)
		return nil, err
	}

	successMsg = "User deleted successfully!"

	return &successMsg, nil
}

// # Update 'Username'
func UpdateUsername(clerkUserID, username string) (*string, error) {
	ctx := context.TODO()

	var err error
	var errMsg string
	var successMsg string

	// # Base 'Filter'
	filter := bson.M{
		"username": username,
	}

	// # Check if the provided 'username' is already 'taken' by the another 'user'
	user, err := GetUser(filter)
	if user != nil {
		if user.ClerkUserID == clerkUserID {
			errMsg = "This is already your current username. Please choose a new username."
		} else {
			errMsg = "Username already taken by another user. Please choose a different username."
		}
		err = errors.New(errMsg)
		return nil, err
	}

	// # Base 'Filter'
	filter = bson.M{
		"clerk_user_id": clerkUserID,
	}

	collection := db.GetMongoCollection(model.UserColl)

	// # Get the 'current system time' in the 'IST' timezone
	time, err := utils.GetCurrentSystemTimeInIST()
	if err != nil {
		utils.LogError(err, "App.UpdateUsername")
		return nil, err
	}

	// # Base 'Set'
	set := bson.M{
		"username":   username,
		"updated_at": time,
		"updated_by": utils.GetActionUser(clerkUserID),
	}

	// # Base 'Update'
	update := bson.M{
		"$set": set,
	}

	// # Update the 'username'
	res, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		utils.LogError(err, "App.UpdateUsername")
		errMsg = "Failed to update the username"
		err = errors.New(errMsg)
		return nil, err
	}
	if res.MatchedCount == 0 {
		errMsg = "User not found"
		err = errors.New(errMsg)
		return nil, err
	}
	if res.ModifiedCount == 0 {
		errMsg = "Username not updated"
		err = errors.New(errMsg)
		return nil, err
	}

	successMsg = "Username updated successfully!"

	return &successMsg, nil
}

// # Get 'User'
func GetUser(filter bson.M) (*model.User, error) {
	ctx := context.TODO()

	var err error
	var errMsg string

	collection := db.GetMongoCollection(model.UserColl)

	var user model.User

	// # Get the 'user'
	err = collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			utils.LogError(err, "App.GetUser")
			errMsg = "User not found"
			err = errors.New(errMsg)
			return nil, err
		} else {
			utils.LogError(err, "App.GetUser")
			errMsg = "Failed to get the user"
			err = errors.New(errMsg)
			return nil, err
		}
	}
	if user.IsDeleted {
		utils.LogError(err, "App.GetUser")
		errMsg = "User is already deleted"
		err = errors.New(errMsg)
		return nil, err
	}

	return &user, nil
}

// # Get 'Users'
func GetUsers() ([]*model.User, error) {
	ctx := context.TODO()

	var err error
	var errMsg string

	collection := db.GetMongoCollection(model.UserColl)

	// # Base 'Filter'
	filter := bson.M{
		"is_deleted": bson.M{
			"$ne": true, // # Exclude the 'deleted' users
		},
	}

	// # Base 'Sort'
	sort := bson.M{
		"created_at": -1, // # Sort the 'documents' (users) by the 'created_at' field in the 'descending' order
	}

	// # Set the 'Find' options
	opts := options.Find()
	opts.SetSort(sort)

	var users []*model.User

	cur, err := collection.Find(ctx, filter, opts)
	if err != nil {
		utils.LogError(err, "App.GetUsers")
		errMsg = "Failed to get the users"
		err = errors.New(errMsg)
		return nil, err
	}

	err = cur.All(ctx, &users)
	if err != nil {
		utils.LogError(err, "App.GetUsers")
		errMsg = "Failed to decode the users"
		err = errors.New(errMsg)
		return nil, err
	}

	return users, nil
}
