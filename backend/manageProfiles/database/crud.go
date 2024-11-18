package database

import (
	"fmt"
	"log"
	"profile/models"
)

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

	// Insert the document
	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		return fmt.Errorf("failed to insert document: %v", err)
	}

	// Print the inserted document's ID
	log.Printf("Document inserted with ID: %v\n", result.InsertedID)
	return nil
}
