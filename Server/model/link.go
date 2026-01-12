package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Link struct {
	ID             *primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID         *primitive.ObjectID `json:"user_id,omitempty" bson:"user_id,omitempty"`
	ClerkUserID    string              `json:"clerk_user_id,omitempty" bson:"clerk_user_id,omitempty"`
	IsUserDeleted  bool                `json:"is_user_deleted,omitempty" bson:"is_user_deleted,omitempty"`
	Title          string              `json:"title,omitempty" bson:"title,omitempty"`
	Description    string              `json:"description,omitempty" bson:"description,omitempty"`
	DestinationURL string              `json:"destination_url,omitempty" bson:"destination_url,omitempty"`
	Slug           string              `json:"slug,omitempty" bson:"slug,omitempty"`
	CreatedAt      *time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	CreatedBy      string              `json:"created_by,omitempty" bson:"created_by,omitempty"`
	UpdatedAt      *time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	UpdatedBy      string              `json:"updated_by,omitempty" bson:"updated_by,omitempty"`
}
