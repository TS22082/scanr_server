// package db provides functionality for interacting with a MongoDB database.

package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

// Connect establishes a connection to a MongoDB instance.
// It initializes a new MongoDB client and attempts to connect to the instance
// specified by the MongoDB URI ("mongodb://localhost:27017").
// The function will log a fatal error and terminate the program if it fails to create
// a new client or to establish a connection to the MongoDB instance.
// Upon successful connection, it prints a confirmation message.
//
// Example usage:
//
//	db.Connect()
func Connect() {
	var err error
	Client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = Client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to MongoDB.")
}
