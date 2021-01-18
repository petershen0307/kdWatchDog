package main

import (
	"context"
	"fmt"
	"log"

	"github.com/petershen0307/kdWatchDog/bot"
	"github.com/petershen0307/kdWatchDog/config"
	"github.com/petershen0307/kdWatchDog/db"
	"github.com/petershen0307/kdWatchDog/models"
	"gopkg.in/mgo.v2/bson"
	tg "gopkg.in/tucnak/telebot.v2"
)

func main() {
	configs := config.Get()
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
	cursor, err = stockColl.Find(context.Background(), bson.M{})
	allStock := []models.StockInfo{}
	cursor.All(context.Background(), &allStock)
	// create stock map
	stockMap := map[string]models.StockInfo{}
	for _, stock := range allStock {
		stockMap[stock.ID] = stock
	}
	bot := bot.New(*configs)
	prefixMsg := "stockID    close    dailyK    dailyD    weeklyK    weeklyD    monthlyK    monthlyD"
	for _, user := range allUsers {
		msg := prefixMsg
		for _, stockID := range user.Stocks {
			msg += fmt.Sprintf("%v    %v    %v    %v    %v    %v    %v    %v\n", stockID,
				stockMap[stockID].DailyPrice.Close,
				stockMap[stockID].DailyKD.K,
				stockMap[stockID].DailyKD.D,
				stockMap[stockID].WeeklyKD.K,
				stockMap[stockID].WeeklyKD.D,
				stockMap[stockID].MonthlyKD.K,
				stockMap[stockID].MonthlyKD.D)
		}
		bot.Send(&tg.User{ID: user.UserID}, msg)
	}
}
