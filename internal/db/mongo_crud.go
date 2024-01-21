package db

import (
	"context"
	db_const "go_server/config/db"
	"go_server/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DbName = db_const.DbName

// GetCollection returns a reference to a MongoDB collection by its name.
func GetCollection(collection string) *mongo.Collection {
	return Client.Database(DbName).Collection(collection)
}

// Create inserts a new document into the specified collection
// that belongs to the model that is passed.
// Returns the result of the insert operation and any error encountered.
func Create(collection string, data interface{}) (*mongo.InsertOneResult, error) {
	c := GetCollection(collection)
	return c.InsertOne(context.Background(), data)
}

// GetAll retrieves all documents from the specified collection that match the provided filter.
// Returns a cursor for the retrieved documents and any error encountered.
func GetAll(collection string, filter interface{}) (*mongo.Cursor, error) {
	c := GetCollection(collection)
	return c.Find(context.Background(), filter)
}

// GetOne retrieves a single document from the specified collection that matches the provided filter.
// Returns the result containing the document or an error if no document is found.
func GetOne(collection string, filter interface{}) *mongo.SingleResult {
	c := GetCollection(collection)
	return c.FindOne(context.Background(), filter)
}

// GetRange retrieves documents from the specified collection that match the provided filter,
// applying skip and limit for pagination. Returns a cursor for the retrieved documents
// and any error encountered.
func GetRange(collection string, filter interface{}, skip int64, limit int64) (*mongo.Cursor, error) {
	c := GetCollection(collection)
	return c.Find(context.Background(), filter, options.Find().SetSkip(skip).SetLimit(limit))
}

// UpdateOne updates a single document in the specified collection based on the provided filter.
// Returns the result of the update operation and any error encountered.
func UpdateOne(collection string, filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	c := GetCollection(collection)
	var bsonUpdate = bson.M{"$set": update}

	return c.UpdateOne(context.Background(), filter, bsonUpdate)
}

// UpdateMany updates multiple documents in the specified collection based on the provided filter.
// Returns the result of the update operation and any error encountered.
func UpdateMany(collection string, filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	c := GetCollection(collection)
	return c.UpdateMany(context.Background(), filter, update)
}

// DeleteOne deletes a single document from the specified collection based on the provided filter.
// Returns the result of the delete operation and any error encountered.
func DeleteOne(collection string, filter interface{}) (*mongo.DeleteResult, error) {
	c := GetCollection(collection)
	return c.DeleteOne(context.Background(), filter)
}

func DeleteOneById(collection string, id string) (*mongo.DeleteResult, error) {
	c := GetCollection(collection)
	objId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objId}
	return c.DeleteOne(context.Background(), filter)
}

// DeleteMany deletes multiple documents from the specified collection based on the provided filter.
// Returns the result of the delete operation and any error encountered.
func DeleteMany(collection string, filter interface{}) (*mongo.DeleteResult, error) {
	c := GetCollection(collection)
	return c.DeleteMany(context.Background(), filter)
}

// GetUserByEmail retrieves a user from the database based on the provided email.
// Returns the user object and any error encountered.
func GetUserByToken(tokenId string) (*models.User, error) {
	var user models.User
	var register_token models.Register_Token

	// Convert tokenId to ObjectID
	objTokenId, err := primitive.ObjectIDFromHex(tokenId)
	if err != nil {
		return nil, err
	}

	// Fetch the registration token
	filter := bson.M{"_id": objTokenId}
	err = GetOne(db_const.Tokens, filter).Decode(&register_token)
	if err != nil {
		return nil, err
	}

	// Fetch the user using UserID from the registration token
	filter = bson.M{"_id": register_token.UserID}
	err = GetOne(db_const.Users, filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// GetUserByEmail retrieves a user from the database based on the provided email.
// Returns the user object and any error encountered.
// this should be depreciated
func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	filter := bson.M{"email": email}
	err := GetOne("users", filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func UserExists(email string) (bool, error) {
	var user models.User

	emailFilter := bson.M{"email": bson.M{"$regex": email, "$options": "i"}}

	filter := bson.M{"$or": []bson.M{emailFilter}}
	err := GetOne("users", filter).Decode(&user)

	// Check if error is because no document was found
	if err == mongo.ErrNoDocuments {
		// User does not exist
		return false, nil
	} else if err != nil {
		// Some other error occurred
		return false, err
	}

	// User exists
	return true, nil
}
