package main

import (
	"context"
	"log"
	"time"

	"github.com/petershen0307/kdWatchDog/config"
	"github.com/petershen0307/kdWatchDog/db"
	"github.com/petershen0307/kdWatchDog/models"
	stockapi "github.com/petershen0307/kdWatchDog/stock-api"
	"github.com/petershen0307/kdWatchDog/stock-api/alphavantage"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	// query watched stocks
	configs := config.Get()
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
	stockIDMap := map[string]models.StockInfo{}
	for _, user := range allUsers {
		for _, id := range user.Stocks {
			if _, ok := stockIDMap[id]; !ok {
				stockIDMap[id] = models.StockInfo{
					ID:         id,
					LastUpdate: time.Date(2021, 1, 1, 1, 1, 1, 0, time.UTC),
				}
			}
		}
	}
	// update to new stock to collection
	stockColl := db.GetCollection(configs.MongoDBURI, configs.DBName, db.CollectionNameStocks)
	for _, stock := range stockIDMap {
		stockColl.UpdateOne(context.Background(),
			bson.M{"stock_id": stock.ID},
			bson.M{"$setOnInsert": stock},
			options.Update().SetUpsert(true))
	}
	// query stock information last update time is greater than one hour
	cursor, err = stockColl.Find(context.Background(), bson.M{"last_update": bson.M{
		"$lt": time.Now().UTC().Add(-1 * time.Hour),
	}})
	if err != nil {
		log.Fatalf("connect to db err=%v", err)
	}
	allStock := []models.StockInfo{}
	cursor.All(context.Background(), &allStock)
	// batch query information
	api := alphavantage.New("...")
	for i, stock := range allStock {
		allStock[i].DailyPrice = api.GetDailyPrice(stock.ID)
		api.Wait()
		allStock[i].DailyKD = api.GetSTOCH(stock.ID, stockapi.Daily, 9, 3, 3, 0, 0)
		api.Wait()
		allStock[i].WeeklyKD = api.GetSTOCH(stock.ID, stockapi.Weekly, 9, 3, 3, 0, 0)
		api.Wait()
		allStock[i].MonthlyKD = api.GetSTOCH(stock.ID, stockapi.Monthly, 9, 3, 3, 0, 0)
		api.Wait()
		allStock[i].LastUpdate = time.Now().UTC()
	}
	// update to DB
	for _, stock := range allStock {
		stockColl.UpdateOne(context.Background(),
			bson.M{"stock_id": stock.ID},
			bson.M{"$set": stock},
			options.Update().SetUpsert(true))
	}
}
