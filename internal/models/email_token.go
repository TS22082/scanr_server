package models

// Email_Token represents the data structure for an emails token in the application.
// It is used to store information about tokens generated during the user registration process.
// The struct includes an ID, the token string, the associated email, and timestamps for creation and last update.
//
// Fields:
//
//	ID: A unique identifier for the token, represented as an ObjectID from MongoDB.
//	Email: The email address associated with the registration token.
//	CreatedAt: A timestamp representing when the token was created.
//	UpdatedAt: A timestamp representing the last time the token information was updated.
//	TokenType: A string representing the type of token (e.g., "registration", "password_reset").

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Email_Token struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Email     string             `bson:"email,omitempty" json:"email"`
	CreatedAt time.Time          `bson:"createdAt,omitempty" json:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt,omitempty" json:"updatedAt"`
	TokenType string             `bson:"tokenType,omitempty" json:"tokenType"`
}
