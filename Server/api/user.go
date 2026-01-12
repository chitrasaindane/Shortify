package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"shortify/app"
	"shortify/schema"
	"shortify/utils"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

// # Handle 'Clerk' User 'Webhook'
func HandleClerkUserWebhook(w http.ResponseWriter, r *http.Request) {
	var err error
	var errMsg string
	var errRes schema.Error

	var res *string

	// # Get the 'clerk user' data from the 'request'
	contentLength := r.ContentLength
	if contentLength == 0 {
		errMsg = "Empty request body provided for the clerk user webhook"
		err = errors.New(errMsg)
		utils.LogError(err, "API.HandleClerkUserWebhook")
		errRes = schema.Error{
			StatusCode: http.StatusBadRequest,
			Message:    errMsg,
		}
		utils.SetAppError(w, &errRes)
		return
	}

	defer r.Body.Close()

	var clerkUserForm schema.ClerkUserRequest

	decoder := json.NewDecoder(r.Body)

	err = decoder.Decode(&clerkUserForm)
	if err != nil {
		utils.LogError(err, "API.HandleClerkUserWebhook")
		errMsg = "Failed to decode the clerk user request body"
		errRes = schema.Error{
			StatusCode: http.StatusBadRequest,
			Message:    errMsg,
		}
		utils.SetAppError(w, &errRes)
		return
	}

	var clerkUser *schema.ClerkUser
	if clerkUserForm.Data == nil {
		errMsg = "Empty clerk user data provided"
		err = errors.New(errMsg)
		utils.LogError(err, "API.HandleClerkUserWebhook")
		errRes = schema.Error{
			StatusCode: http.StatusBadRequest,
			Message:    errMsg,
		}
		utils.SetAppError(w, &errRes)
		return
	} else {
		clerkUser = clerkUserForm.Data
	}

	// # Validate the 'clerk user' data
	if clerkUserForm.Type == "user.deleted" {
		if !clerkUser.Deleted || clerkUser.Object != "user" {
			errMsg = "Invalid delete user request found"
			err = errors.New(errMsg)
			utils.LogError(err, "API.HandleClerkUserWebhook")
			errRes = schema.Error{
				StatusCode: http.StatusBadRequest,
				Message:    errMsg,
			}
			utils.SetAppError(w, &errRes)
			return
		}
	} else {
		if clerkUser.FirstName == nil || utils.GetStringValue(clerkUser.FirstName) == "" {
			errMsg = "User first name cannot be empty"
			err = errors.New(errMsg)
			utils.LogError(err, "API.HandleClerkUserWebhook")
			errRes = schema.Error{
				StatusCode: http.StatusBadRequest,
				Message:    errMsg,
			}
			utils.SetAppError(w, &errRes)
			return
		}
		if clerkUser.LastName == nil || utils.GetStringValue(clerkUser.LastName) == "" {
			errMsg = "User last name cannot be empty"
			err = errors.New(errMsg)
			utils.LogError(err, "API.HandleClerkUserWebhook")
			errRes = schema.Error{
				StatusCode: http.StatusBadRequest,
				Message:    errMsg,
			}
			utils.SetAppError(w, &errRes)
			return
		}
		if clerkUser.EmailAddresss == nil || clerkUser.EmailAddresss[0] == nil || clerkUser.EmailAddresss[0].EmailAddress == nil || utils.GetStringValue(clerkUser.EmailAddresss[0].EmailAddress) == "" {
			errMsg = "User email address cannot be empty"
			err = errors.New(errMsg)
			utils.LogError(err, "API.HandleClerkUserWebhook")
			errRes = schema.Error{
				StatusCode: http.StatusBadRequest,
				Message:    errMsg,
			}
			utils.SetAppError(w, &errRes)
			return
		}
	}

	switch clerkUserForm.Type {
	case "user.created":
		{
			// # Create the 'user'
			res, err = app.CreateUser(clerkUser)
		}
	case "user.updated":
		{
			// # Update the 'user'
			res, err = app.UpdateUser(clerkUser)
		}
	case "user.deleted":
		{
			// # Delete the 'user'
			res, err = app.DeleteUser(clerkUser)
		}
	default:
		{
			errMsg = fmt.Sprintf("Invalid clerk user webhook type '%s' provided", clerkUserForm.Type)
			err = errors.New(errMsg)
		}
	}

	if err != nil {
		utils.LogError(err, "API.HandleClerkUserWebhook")
		errMsg = err.Error()
		errRes = schema.Error{
			StatusCode: http.StatusInternalServerError,
			Message:    errMsg,
		}
		utils.SetAppError(w, &errRes)
		return
	}

	utils.SetAppResponse(w, res)
}

// # Update 'Username'
func UpdateUsername(w http.ResponseWriter, r *http.Request) {
	var err error
	var errMsg string
	var errRes schema.Error

	// # Get the 'clerk user ID' from the 'query params'
	clerkUserID := r.URL.Query().Get("clerk_user_id")
	clerkUserID = utils.GetTrimmedValue(clerkUserID)
	if clerkUserID == "" {
		errMsg = "Empty clerk user ID provided"
		errRes = schema.Error{
			StatusCode: http.StatusBadRequest,
			Message:    errMsg,
		}
		utils.SetAppError(w, &errRes)
		return
	}

	// # Get the 'username' from the 'query params'
	username := r.URL.Query().Get("username")
	username = utils.GetTrimmedValue(username)
	username = strings.ToLower(username)
	if username == "" {
		errMsg = "Empty username provided"
		errRes = schema.Error{
			StatusCode: http.StatusBadRequest,
			Message:    errMsg,
		}
		utils.SetAppError(w, &errRes)
		return
	}

	// # Update the 'username'
	res, err := app.UpdateUsername(clerkUserID, username)
	if err != nil {
		errMsg = err.Error()
		errRes = schema.Error{
			StatusCode: http.StatusInternalServerError,
			Message:    errMsg,
		}
		utils.SetAppError(w, &errRes)
		return
	}

	utils.SetAppResponse(w, res)
}

// # Get 'User'
func GetUser(w http.ResponseWriter, r *http.Request) {
	var err error
	var errMsg string
	var errRes schema.Error

	// # Get the 'clerk user ID' from the 'query params'
	clerkUserID := r.URL.Query().Get("clerk_user_id")
	clerkUserID = utils.GetTrimmedValue(clerkUserID)
	if clerkUserID == "" {
		errMsg = "Empty clerk user ID provided"
		errRes = schema.Error{
			StatusCode: http.StatusBadRequest,
			Message:    errMsg,
		}
		utils.SetAppError(w, &errRes)
		return
	}

	// # Base 'Filter'
	filter := bson.M{
		"clerk_user_id": clerkUserID,
	}

	// # Get the 'user'
	user, err := app.GetUser(filter)
	if err != nil {
		errMsg = err.Error()
		errRes = schema.Error{
			StatusCode: http.StatusInternalServerError,
			Message:    errMsg,
		}
		utils.SetAppError(w, &errRes)
		return
	}

	utils.SetAppResponse(w, user)
}

// # Get 'Users'
func GetUsers(w http.ResponseWriter, r *http.Request) {
	var err error
	var errMsg string
	var errRes schema.Error

	// # Get the 'users'
	users, err := app.GetUsers()
	if err != nil {
		errMsg = err.Error()
		errRes = schema.Error{
			StatusCode: http.StatusInternalServerError,
			Message:    errMsg,
		}
		utils.SetAppError(w, &errRes)
		return
	}

	utils.SetAppResponse(w, users)
}
