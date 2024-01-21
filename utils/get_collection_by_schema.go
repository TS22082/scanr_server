package utils

import (
	"reflect"
)

// GetCollectionBySchema returns the name of the MongoDB collection based on the provided schema type.
//
// The function takes a reflect.Type parameter representing the schema type and
// returns the corresponding collection name as a string.
//
// Example usage:
//
//	collectionName := GetCollectionBySchema(reflect.TypeOf(User{}))
//	// collectionName will be "users" if User is the schema type.
//
// Supported schema types:
// - "User" returns "users"
// - "Register_Token" returns "register_tokens"
//
// If the provided schema type is not supported, an empty string is returned.
func GetCollectionBySchema(schema reflect.Type) string {

	typeName := schema.Name()

	switch typeName {
	case "User":
		return "users"
	case "Register_Token":
		return "register_tokens"

	default:
		return ""
	}
}
