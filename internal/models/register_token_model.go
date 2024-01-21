// Package models defines the data structures and models used throughout the application.
// This includes definitions for database schemas and any other data-related structures.
package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Register_Token represents the data structure for a registration token in the application.
// It is used to store information about tokens generated during the user registration process.
// The struct includes an ID, the token string, the associated email, and timestamps for creation and last update.
//
// Fields:
//
//	ID: A unique identifier for the token, represented as an ObjectID from MongoDB.
//	Email: The email address associated with the registration token.
//	CreatedAt: A timestamp representing when the token was created.
//	UpdatedAt: A timestamp representing the last time the token information was updated.
type Register_Token struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID    primitive.ObjectID `bson:"userId,omitempty" json:"userId"`
	CreatedAt time.Time          `bson:"createdAt,omitempty" json:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt,omitempty" json:"updatedAt"`
}
