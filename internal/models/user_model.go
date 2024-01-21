// Package models contains the data structures and models used in the application.
// These models define the structure of data objects and their representation in the database.
package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User represents a user in the system.
// Each user has a unique ID, along with associated personal information like username, email, and password.
// The User struct also tracks the creation and last update times.
//
// Fields:
//
//	ID: A unique identifier for the user, represented as a MongoDB ObjectID.
//	Username: The user's chosen username. This field is indexed in MongoDB for efficient lookups.
//	Email: The user's email address. Like the Username, this is also indexed.
//	Password: The user's password. This is stored as a hash for security purposes.
//	CreatedAt: The timestamp of when the user account was created.
//	UpdatedAt: The timestamp of the last update made to the user's information.
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username  string             `bson:"username" json:"username"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"password"`
	Verified  bool               `bson:"verified" json:"verified"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt" json:"updatedAt"`
}
