package main

import (
	"context"
	"log"

	"github.com/petershen0307/kdWatchDog/bot"
	"github.com/petershen0307/kdWatchDog/config"
	"github.com/petershen0307/kdWatchDog/db"
	"github.com/petershen0307/kdWatchDog/handlers"
	"github.com/petershen0307/kdWatchDog/models"
	"gopkg.in/mgo.v2/bson"
	tg "gopkg.in/tucnak/telebot.v2"
)

func main() {
	configs := config.Get()

	// query all users
	userColl := db.GetCollection(configs.MongoDBURI, configs.DBName, "users")
	cursor, err := userColl.Find(
		context.Background(),
		bson.M{},
	)
	if err != nil {
		log.Fatalf("connect to db err=%v", err)
	}
	allUsers := []models.User{}
	cursor.All(context.Background(), &allUsers)
	stockColl := db.GetCollection(configs.MongoDBURI, configs.DBName, "stocks")

	// stock map
	stockMap := handlers.GetStockMap(stockColl)
	bot := bot.New(*configs)
	for _, user := range allUsers {
		msg := handlers.RenderOneUserOutput(&user, stockMap)
		bot.SendAlbum(&tg.User{ID: user.UserID}, tg.Album{&tg.Photo{File: tg.FromReader(msg)}})
	}
}
