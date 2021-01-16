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

const addCommand = "/add"

func getAddStockHandler(responseCallback responseCallbackFunc, collection *mongo.Collection) (string, func(*tg.Message)) {
	return addCommand, func(m *tg.Message) {
		stockID := strings.Replace(m.Text, addCommand, "", 1)
		stockID = strings.TrimSpace(stockID)
		stockID = strings.ToUpper(stockID)
		var responseMsg *string = new(string)
		*responseMsg = fmt.Sprintf("add %v ok", stockID)

		defer responseCallback(m.Sender, responseMsg)

		if err := stockIDValidator(stockID); err != nil {
			log.Printf("user(%v) insert invalid stock id(%v) err = %v", m.Sender.ID, stockID, err)
			*responseMsg = "invalid stock id"
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
				*responseMsg = fmt.Sprintf("Add stock (%v) with err = %v", stockID, err)
				log.Printf("Add stock for user(%v) with err = %v", m.Sender.ID, err)
				return
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
			user.LastUpdate = time.Now().UTC()
			_, err = collection.UpdateOne(context.Background(), bson.M{"user_id": m.Sender.ID}, bson.M{"$set": user}, options.Update().SetUpsert(true))
			if err != nil {
				*responseMsg = fmt.Sprintf("Add stock fail (%v)", err)
			}
		}
	}
}

func stockIDValidator(stockID string) error {
	// [ToDo] query real stock id from tw and us
	if stockID == "" {
		return fmt.Errorf("stock id is empty")
	}
	return nil
}
