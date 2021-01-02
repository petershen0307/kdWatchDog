package main

import (
	"context"
	"log"

	"github.com/petershen0307/kdWatchDog/bot"
	"github.com/petershen0307/kdWatchDog/config"
	"github.com/petershen0307/kdWatchDog/handlers"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	configs := config.Get()

	tgBot := bot.New(*configs)

	// Set client options
	clientOptions := options.Client().ApplyURI(configs.MongoDBURI)
	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	collection := client.Database(configs.DBName).Collection("users")

	handlers.RegisterHandlers(tgBot, collection)
	tgBot.Start()
}
