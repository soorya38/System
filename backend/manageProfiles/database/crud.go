package database

import (
	"fmt"
	"log"
	"profile/models"

	"go.mongodb.org/mongo-driver/bson"
)

// CreateProfile creates a profile int the MongoDB collection by its ID.
func CreateProfile(user models.Profile) error {
	uri := "mongodb+srv://hackathon:hackathon@hackathon.1yrkb.mongodb.net/?retryWrites=true&w=majority&appName=hackathon"
	db_name := "management_system"
	collection_name := "users"

	client, ctx, err := ConnectMongoDB(uri)
	if err != nil {
		return fmt.Errorf("error : %v", err)
	}

	db := client.Database(db_name)
	collection := db.Collection(collection_name)

	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		return fmt.Errorf("failed to insert document: %v", err)
	}

	log.Printf("Document inserted with ID: %v\n", result.InsertedID)
	return nil
}

// DeleteProfile deletes a profile from the MongoDB collection by its ID.
func DeleteProfile(id string) error {
	uri := "mongodb+srv://hackathon:hackathon@hackathon.1yrkb.mongodb.net/?retryWrites=true&w=majority&appName=hackathon"
	dbName := "management_system"
	collectionName := "users"

	// Connect to MongoDB
	client, ctx, err := ConnectMongoDB(uri)
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}
	defer client.Disconnect(ctx) // Ensure the connection is closed when done

	db := client.Database(dbName)
	collection := db.Collection(collectionName)

	// Delete the document with the specified ID
	filter := bson.M{"id": id}
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to delete document: %v", err)
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("no document found with the ID: %v", id)
	}

	log.Printf("Document deleted with ID: %v\n", id)
	return nil
}
