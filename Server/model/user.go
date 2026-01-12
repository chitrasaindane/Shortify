package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          *primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ClerkUserID string              `json:"clerk_user_id,omitempty" bson:"clerk_user_id,omitempty"`
	Code        int64               `json:"code,omitempty" bson:"code,omitempty"`
	FirstName   string              `json:"first_name,omitempty" bson:"first_name,omitempty"`
	LastName    string              `json:"last_name,omitempty" bson:"last_name,omitempty"`
	Email       string              `json:"email,omitempty" bson:"email,omitempty"`
	Username    string              `json:"username,omitempty" bson:"username,omitempty"`
	CreatedAt   *time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	CreatedBy   string              `json:"created_by,omitempty" bson:"created_by,omitempty"`
	UpdatedAt   *time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	UpdatedBy   string              `json:"updated_by,omitempty" bson:"updated_by,omitempty"`
	IsDeleted   bool                `json:"is_deleted,omitempty" bson:"is_deleted,omitempty"`
	DeletedAt   *time.Time          `json:"deleted_at,omitempty" bson:"deleted_at,omitempty"`
	DeletedBy   string              `json:"deleted_by,omitempty" bson:"deleted_by,omitempty"`
}
