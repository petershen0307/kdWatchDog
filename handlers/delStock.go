package handlers

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/petershen0307/kdWatchDog/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
	tg "gopkg.in/tucnak/telebot.v2"
)

const delCommand = "/del"

func getDelStockHandler(responseCallback responseCallbackFunc, collection *mongo.Collection) (string, func(*tg.Message)) {
	return delCommand, func(m *tg.Message) {
		stockID := strings.Replace(m.Text, delCommand, "", 1)
		stockID = strings.TrimSpace(stockID)
		var p *post = &post{
			to:   m.Sender,
			what: "success",
		}
		defer responseCallback(p)

		queryUser := models.User{}
		err := collection.FindOne(
			context.Background(),
			bson.M{"user_id": m.Sender.ID},
		).Decode(&queryUser)
		if err == mongo.ErrNoDocuments {
			p.what = "no record"
			return
		}
		if err != nil {
			p.what = err.Error()
			log.Printf("Query user(%v) with err = %v", m.Sender.ID, err)
			return
		}

		removeIndex := -1
		for i, v := range queryUser.Stocks {
			if v == stockID {
				removeIndex = i
				break
			}
		}
		if removeIndex != -1 {
			queryUser.Stocks = append(queryUser.Stocks[0:removeIndex], queryUser.Stocks[removeIndex+1:]...)
		}
		queryUser.LastUpdate = time.Now().UTC()
		_, err = collection.UpdateOne(context.Background(), bson.M{"user_id": m.Sender.ID}, bson.M{"$set": queryUser},
			options.Update().SetUpsert(true))
		if err != nil {
			p.what = fmt.Sprintf("Delete stock(%v) fail (%v)", stockID, err)
		}
	}
}
