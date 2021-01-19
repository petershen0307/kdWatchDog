package handlers

import (
	"context"
	"fmt"
	"log"

	"github.com/petershen0307/kdWatchDog/models"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
	tg "gopkg.in/tucnak/telebot.v2"
)

const queryCommand = "/query"

func getQueryStockHandler(responseCallback responseCallbackFunc, userColl, stockColl *mongo.Collection) (string, func(*tg.Message)) {
	return queryCommand, func(m *tg.Message) {
		var responseMsg *string = new(string)

		defer responseCallback(m.Sender, responseMsg)

		// query all user
		var user models.User
		if err := userColl.FindOne(context.Background(), bson.M{"user_id": m.Sender.ID}).Decode(&user); err != nil {
			log.Printf("connect to db err=%v", err)
			*responseMsg = "connect user collection failed"
			return
		}

		stockMap := GetStockMap(stockColl)
		if stockMap == nil {
			*responseMsg = "connect user collection failed"
			return
		}
		*responseMsg = RenderOneUserOutput(&user, stockMap)
	}
}

// GetStockMap create stock map cache
func GetStockMap(stockColl *mongo.Collection) map[string]models.StockInfo {
	// query all stock info
	cursor, err := stockColl.Find(context.Background(), bson.M{})
	if err != nil {
		log.Printf("connect to db err=%v", err)
		return nil
	}
	allStock := []models.StockInfo{}
	cursor.All(context.Background(), &allStock)

	// create stock map
	stockMap := map[string]models.StockInfo{}
	for _, stock := range allStock {
		stockMap[stock.ID] = stock
	}
	return stockMap
}

// RenderOneUserOutput render one user stock info
func RenderOneUserOutput(user *models.User, stockMap map[string]models.StockInfo) string {
	responseMsg := "stockID    close    dailyK    dailyD    weeklyK    weeklyD    monthlyK    monthlyD\n"
	for _, stockID := range user.Stocks {
		responseMsg += fmt.Sprintf("%v    %v    %v    %v    %v    %v    %v    %v\n", stockID,
			stockMap[stockID].DailyPrice.Close,
			stockMap[stockID].DailyKD.K,
			stockMap[stockID].DailyKD.D,
			stockMap[stockID].WeeklyKD.K,
			stockMap[stockID].WeeklyKD.D,
			stockMap[stockID].MonthlyKD.K,
			stockMap[stockID].MonthlyKD.D)
	}
	return responseMsg
}
