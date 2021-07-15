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
)

const delCommand = "/del"

// DelStock handle delete command
func (handle *Handler) DelStock(mail *Mail) {
	stockID := strings.Replace(mail.fromMsg, delCommand, "", 1)
	stockID = strings.TrimSpace(stockID)
	defer func() {
		handle.mailbox <- *mail
	}()
	mail.toMsg = "success"
	queryUser := models.User{}
	err := handle.userColl.FindOne(
		context.Background(),
		bson.M{"user_id": mail.userID},
	).Decode(&queryUser)
	if err == mongo.ErrNoDocuments {
		mail.toMsg = "no record"
		return
	}
	if err != nil {
		mail.toMsg = err.Error()
		log.Printf("Query user(%v) with err = %v", mail.userID, err)
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
	_, err = handle.userColl.UpdateOne(context.Background(), bson.M{"user_id": mail.userID}, bson.M{"$set": queryUser},
		options.Update().SetUpsert(true))
	if err != nil {
		mail.toMsg = fmt.Sprintf("Delete stock(%v) fail (%v)", stockID, err)
	}
}
