package handlers

import (
	"context"
	"fmt"
	"strings"

	"github.com/petershen0307/kdWatchDog/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
	tg "gopkg.in/tucnak/telebot.v2"
)

var addCommand = "/add"

func getAddStockHandler(responseCallback func(tg.Recipient, interface{}, ...interface{}), collection *mongo.Collection) (string, func(*tg.Message)) {
	return addCommand, func(m *tg.Message) {
		stockID := strings.Replace(m.Text, addCommand, "", 1)
		stockID = strings.TrimSpace(stockID)
		responseMsg := fmt.Sprintf("add %v ok", stockID)
		defer responseCallback(
			m.Sender,
			responseMsg,
		)

		if stockID == "" {
			responseMsg = "stock id is empty"
			return
		}
		var user models.User
		err := collection.FindOne(context.Background(), bson.M{"user_id": m.Sender.ID}).Decode(&user)
		updateDB := false
		if err != nil {
			// insert
			if err == mongo.ErrNoDocuments {
				updateDB = true
			} else {
				responseMsg = fmt.Sprintf("Add stock fail (%v)", err)
			}
		} else {
			// update
			updateDB = true
			for _, stock := range user.Stocks {
				if stock == stockID {
					updateDB = false
					break
				}
			}
		}
		if updateDB {
			user.UserID = m.Sender.ID
			user.Stocks = append(user.Stocks, stockID)
			_, err = collection.UpdateOne(context.Background(), bson.M{"user_id": m.Sender.ID}, bson.M{"$set": user}, options.Update().SetUpsert(true))
			if err != nil {
				responseMsg = fmt.Sprintf("Add stock fail (%v)", err)
			}
		}
	}
}
