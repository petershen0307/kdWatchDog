package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetCollection to get mongoDB collection
func GetCollection(mongoDBURI, dbName, collectionName string) *mongo.Collection {
	// Set client options
	clientOptions := options.Client().ApplyURI(mongoDBURI)
	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	collection := client.Database(dbName).Collection(collectionName)
	return collection
}
