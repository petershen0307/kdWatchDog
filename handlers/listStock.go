package handlers

import (
	"context"
	"log"
	"sort"
	"strings"

	"github.com/petershen0307/kdWatchDog/models"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
	tg "gopkg.in/tucnak/telebot.v2"
)

const listCommand = "/list"

func getListStockHandler(responseCallback responseCallbackFunc, collection *mongo.Collection) (string, botHandler) {
	return listCommand, func(m *tg.Message) {
		var responseMsg *string = new(string)
		*responseMsg = "no data"

		defer responseCallback(m.Sender, responseMsg)

		queryUser := models.User{}
		err := collection.FindOne(
			context.Background(),
			bson.M{"user_id": m.Sender.ID},
		).Decode(&queryUser)
		if err == mongo.ErrNoDocuments {
			return
		}
		if err != nil {
			*responseMsg = err.Error()
			log.Printf("Query user(%v) with err = %v", m.Sender.ID, err)
			return
		}

		// [ToDo] add stock name and region(tw, us)
		// Region ID Name
		sort.Strings(queryUser.Stocks)
		*responseMsg = strings.Join(queryUser.Stocks, "\n")
	}
}
