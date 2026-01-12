package api

import (
	"encoding/json"
	"net/http"
	"shortify/app"
	"shortify/schema"
	"shortify/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// # Create 'Link'
func CreateLink(w http.ResponseWriter, r *http.Request) {
	var err error
	var errMsg string
	var errRes schema.Error

	// # Get the 'link' data from the 'request'
	contentLength := r.ContentLength
	if contentLength == 0 {
		errMsg = "Empty request body provided for the create link"
		errRes = schema.Error{
			StatusCode: http.StatusBadRequest,
			Message:    errMsg,
		}
		utils.SetAppError(w, &errRes)
		return
	}

	defer r.Body.Close()

	var linkForm schema.LinkRequest

	decoder := json.NewDecoder(r.Body)

	err = decoder.Decode(&linkForm)
	if err != nil {
		utils.LogError(err, "API.CreateLink")
		errMsg = "Failed to decode the create link request body"
		errRes = schema.Error{
			StatusCode: http.StatusBadRequest,
			Message:    errMsg,
		}
		utils.SetAppError(w, &errRes)
		return
	}

	var link *schema.Link
	if linkForm.Data == nil {
		errMsg = "Empty link data provided"
		errRes = schema.Error{
			StatusCode: http.StatusBadRequest,
			Message:    errMsg,
		}
		utils.SetAppError(w, &errRes)
		return
	} else {
		link = linkForm.Data
	}

	// # Validate the 'link' data
	if link.DestinationURL == nil || utils.GetStringValue(link.DestinationURL) == "" {
		errMsg = "Link destination URL cannot be empty"
		errRes = schema.Error{
			StatusCode: http.StatusBadRequest,
			Message:    errMsg,
		}
		utils.SetAppError(w, &errRes)
		return
	}
	if link.Slug == nil || utils.GetStringValue(link.Slug) == "" {
		errMsg = "Link slug cannot be empty"
		errRes = schema.Error{
			StatusCode: http.StatusBadRequest,
			Message:    errMsg,
		}
		utils.SetAppError(w, &errRes)
		return
	}

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

	// # Create the 'link'
	res, err := app.CreateLink(clerkUserID, link)
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

// # Update 'Link'
func UpdateLink(w http.ResponseWriter, r *http.Request) {
	var err error
	var errMsg string
	var errRes schema.Error

	// # Get the 'link' data from the 'request'
	contentLength := r.ContentLength
	if contentLength == 0 {
		errMsg = "Empty request body provided for the update link"
		errRes = schema.Error{
			StatusCode: http.StatusBadRequest,
			Message:    errMsg,
		}
		utils.SetAppError(w, &errRes)
		return
	}

	defer r.Body.Close()

	var linkForm schema.LinkRequest

	decoder := json.NewDecoder(r.Body)

	err = decoder.Decode(&linkForm)
	if err != nil {
		utils.LogError(err, "API.UpdateLink")
		errMsg = "Failed to decode the update link request body"
		errRes = schema.Error{
			StatusCode: http.StatusBadRequest,
			Message:    errMsg,
		}
		utils.SetAppError(w, &errRes)
		return
	}

	var link *schema.Link
	if linkForm.Data == nil {
		errMsg = "Empty link data provided"
		errRes = schema.Error{
			StatusCode: http.StatusBadRequest,
			Message:    errMsg,
		}
		utils.SetAppError(w, &errRes)
		return
	} else {
		link = linkForm.Data
	}

	// # Validate the 'link' data
	if link.DestinationURL == nil || utils.GetStringValue(link.DestinationURL) == "" {
		errMsg = "Link destination URL cannot be empty"
		errRes = schema.Error{
			StatusCode: http.StatusBadRequest,
			Message:    errMsg,
		}
		utils.SetAppError(w, &errRes)
		return
	}
	if link.Slug == nil || utils.GetStringValue(link.Slug) == "" {
		errMsg = "Link slug cannot be empty"
		errRes = schema.Error{
			StatusCode: http.StatusBadRequest,
			Message:    errMsg,
		}
		utils.SetAppError(w, &errRes)
		return
	}

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

	// # Get the 'link ID' from the 'query params'
	linkIDStr := r.URL.Query().Get("link_id")
	linkIDStr = utils.GetTrimmedValue(linkIDStr)
	if linkIDStr == "" {
		errMsg = "Empty link ID provided"
		errRes = schema.Error{
			StatusCode: http.StatusBadRequest,
			Message:    errMsg,
		}
		utils.SetAppError(w, &errRes)
		return
	}

	// # Convert the 'link ID' from type 'string' to type 'object ID'
	linkID, err := primitive.ObjectIDFromHex(linkIDStr)
	if err != nil {
		utils.LogError(err, "API.DeleteLink")
		errMsg = "Invalid link ID provided"
		errRes = schema.Error{
			StatusCode: http.StatusBadRequest,
			Message:    errMsg,
		}
		utils.SetAppError(w, &errRes)
		return
	}

	// # Update the 'link'
	res, err := app.UpdateLink(clerkUserID, &linkID, link)
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

// # Delete 'Link'
func DeleteLink(w http.ResponseWriter, r *http.Request) {
	var err error
	var errMsg string
	var errRes schema.Error

	// # Get the 'link ID' from the 'query params'
	linkIDStr := r.URL.Query().Get("link_id")
	linkIDStr = utils.GetTrimmedValue(linkIDStr)
	if linkIDStr == "" {
		errMsg = "Empty link ID provided"
		errRes = schema.Error{
			StatusCode: http.StatusBadRequest,
			Message:    errMsg,
		}
		utils.SetAppError(w, &errRes)
		return
	}

	// # Convert the 'link ID' from type 'string' to type 'object ID'
	linkID, err := primitive.ObjectIDFromHex(linkIDStr)
	if err != nil {
		utils.LogError(err, "API.DeleteLink")
		errMsg = "Invalid link ID provided"
		errRes = schema.Error{
			StatusCode: http.StatusBadRequest,
			Message:    errMsg,
		}
		utils.SetAppError(w, &errRes)
		return
	}

	// # Delete the 'link'
	res, err := app.DeleteLink(&linkID)
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

// # Get 'Link'
func GetLink(w http.ResponseWriter, r *http.Request) {
	var err error
	var errMsg string
	var errRes schema.Error

	// # Get the 'link ID' from the 'query params'
	linkIDStr := r.URL.Query().Get("link_id")
	linkIDStr = utils.GetTrimmedValue(linkIDStr)
	if linkIDStr == "" {
		errMsg = "Empty link ID provided"
		errRes = schema.Error{
			StatusCode: http.StatusBadRequest,
			Message:    errMsg,
		}
		utils.SetAppError(w, &errRes)
		return
	}

	// # Convert the 'link ID' from type 'string' to type 'object ID'
	linkID, err := primitive.ObjectIDFromHex(linkIDStr)
	if err != nil {
		utils.LogError(err, "API.DeleteLink")
		errMsg = "Invalid link ID provided"
		errRes = schema.Error{
			StatusCode: http.StatusBadRequest,
			Message:    errMsg,
		}
		utils.SetAppError(w, &errRes)
		return
	}

	// # Base 'Filter'
	filter := bson.M{
		"_id": linkID,
	}

	// # Get the 'link'
	res, err := app.GetLink(filter)
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

// # Get 'Links'
func GetLinks(w http.ResponseWriter, r *http.Request) {
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

	// # Get the 'links'
	res, err := app.GetLinks(clerkUserID)
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
