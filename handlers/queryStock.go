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
		var p *post = &post{
			to:      m.Sender,
			what:    "no data",
			options: []interface{}{},
		}
		defer responseCallback(p)

		// query all user
		var user models.User
		if err := userColl.FindOne(context.Background(), bson.M{"user_id": m.Sender.ID}).Decode(&user); err != nil {
			log.Printf("connect to db err=%v", err)
			p.what = "connect user collection failed"
			return
		}

		stockMap := GetStockMap(stockColl)
		if stockMap == nil {
			p.what = "connect user collection failed"
			return
		}
		msg, tgParseMode := RenderOneUserOutput(&user, stockMap)
		p.what = msg
		p.options = append(p.options, tgParseMode)
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
func RenderOneUserOutput(user *models.User, stockMap map[string]models.StockInfo) (string, tg.ParseMode) {
	responseMsg := fmt.Sprint(
		"| stockID   | close     | dayK      | dayD      | weekK     | weekD | monthK | monthD |\n",
		"|:----------|:----------|:----------|:----------|:----------|:------|:-------|:-------|\n",
	)

	for _, stockID := range user.Stocks {
		responseMsg += fmt.Sprintf("| %v | %v | %v | %v | %v | %v | %v | %v |\n",
			stockID,
			stockMap[stockID].DailyPrice.Close,
			stockMap[stockID].DailyKD.K,
			stockMap[stockID].DailyKD.D,
			stockMap[stockID].WeeklyKD.K,
			stockMap[stockID].WeeklyKD.D,
			stockMap[stockID].MonthlyKD.K,
			stockMap[stockID].MonthlyKD.D)
	}
	return responseMsg, tg.ModeMarkdown
}
