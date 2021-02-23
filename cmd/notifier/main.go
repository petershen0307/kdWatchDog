package main

import (
	"context"
	"log"

	"github.com/petershen0307/kdWatchDog/bot"
	"github.com/petershen0307/kdWatchDog/config"
	"github.com/petershen0307/kdWatchDog/db"
	"github.com/petershen0307/kdWatchDog/handlers"
	"github.com/petershen0307/kdWatchDog/imgur"
	"github.com/petershen0307/kdWatchDog/models"
	"gopkg.in/mgo.v2/bson"
	tg "gopkg.in/tucnak/telebot.v2"
)

func main() {
	configs := config.Get()

	// query all users
	userColl := db.GetCollection(configs.MongoDBURI, configs.DBName, db.CollectionNameUsers)
	cursor, err := userColl.Find(
		context.Background(),
		bson.M{},
	)
	if err != nil {
		log.Fatalf("connect to db err=%v", err)
	}
	allUsers := []models.User{}
	cursor.All(context.Background(), &allUsers)
	stockColl := db.GetCollection(configs.MongoDBURI, configs.DBName, db.CollectionNameStocks)

	// stock map
	stockMap := handlers.GetStockMap(stockColl)
	bot := bot.New(*configs)
	for _, user := range allUsers {
		msg := handlers.RenderOneUserOutput(&user, stockMap)
		// upload to imgur
		link, err := imgur.UploadImage(configs.ImgurClientID, msg.Bytes())
		if err != nil {
			log.Fatalf("upload image with error=%v", err)
		}
		bot.Send(&tg.User{ID: user.UserID}, link)
	}
}
