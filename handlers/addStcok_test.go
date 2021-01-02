package handlers

import (
	"context"
	"testing"

	"github.com/petershen0307/kdWatchDog/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	tg "gopkg.in/tucnak/telebot.v2"
)

func Test_getAddStockHandler(t *testing.T) {
	t.Skip()
	configs := config.Config{
		MongoDBURI: "mongodb://localhost:27017/ut",
		DBName:     "kdWatchDog",
	}

	clientOptions := options.Client().ApplyURI(configs.MongoDBURI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		t.Fatal(err)
	}
	collection := client.Database(configs.DBName).Collection("users")
	responseCallback := func(to tg.Recipient, what interface{}, options ...interface{}) {
		if what != "add 1234ok" {
			t.Fatalf("Got wrong msg (%v)", what)
		}
	}
	command, f := getAddStockHandler(responseCallback, collection)
	if command != addCommand {
		t.Fatal("Wrong command")
	}
	f(&tg.Message{
		Sender: &tg.User{
			ID: 5566,
		},
		Text: "/add 1234",
	})
}
