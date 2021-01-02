package test

import (
	"context"

	"github.com/petershen0307/kdWatchDog/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

var configs config.Config = config.Config{
	MongoDBURI: "mongodb://localhost:27017/ut",
	DBName:     "ut",
}

// InitDB to get unit test mongoDB client
func InitDB() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(configs.MongoDBURI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// GetCollection to get unit test mongoDB collection
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database(configs.DBName).Collection(collectionName)
}

// Deinit drop ut database
func Deinit(client *mongo.Client, collectionName string) error {
	client.Database(configs.DBName).Collection(collectionName).Drop(context.Background())
	err := client.Disconnect(context.Background())
	return err
}

// RemoveDocs remove all docs
func RemoveDocs(collection *mongo.Collection) error {
	_, err := collection.DeleteMany(context.Background(), bson.M{})
	return err
}
