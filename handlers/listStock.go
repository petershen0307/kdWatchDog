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

func getListStockHandler(responseCallback responseCallbackFunc, collection *mongo.Collection) (string, func(*tg.Message)) {
	return listCommand, func(m *tg.Message) {
		var p *post = &post{
			to:   m.Sender,
			what: "no data",
		}
		defer responseCallback(p)

		queryUser := models.User{}
		err := collection.FindOne(
			context.Background(),
			bson.M{"user_id": m.Sender.ID},
		).Decode(&queryUser)
		if err == mongo.ErrNoDocuments {
			return
		}
		if err != nil {
			p.what = err.Error()
			log.Printf("Query user(%v) with err = %v", m.Sender.ID, err)
			return
		}

		// [ToDo] add stock name and region(tw, us)
		// Region ID Name
		sort.Strings(queryUser.Stocks)
		p.what = strings.Join(queryUser.Stocks, "\n")
	}
}
