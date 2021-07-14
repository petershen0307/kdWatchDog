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

const addCommand = "/add"

func stockIDValidator(stockID string) error {
	// [ToDo] query real stock id from tw and us
	if stockID == "" {
		return fmt.Errorf("stock id is empty")
	}
	return nil
}

// AddStock handle add command
func (handle *Handler) AddStock(mail *Mail) {
	stockID := strings.Replace(mail.fromMsg, addCommand, "", 1)
	stockID = strings.TrimSpace(stockID)
	stockID = strings.ToUpper(stockID)
	responseMail := *mail

	if err := stockIDValidator(stockID); err != nil {
		log.Printf("user(%v) insert invalid stock id(%v) err = %v", mail.userID, stockID, err)
		responseMail.toMsg = "invalid stock id"
		return
	}

	var user models.User
	err := handle.userColl.FindOne(context.Background(), bson.M{"user_id": mail.userID, "bot_platform": mail.platform}).Decode(&user)
	updateDB := false
	if err != nil {
		// insert
		if err == mongo.ErrNoDocuments {
			updateDB = true
		} else {
			responseMail.toMsg = fmt.Sprintf("Add stock (%v) with err = %v", stockID, err)
			log.Printf("Add stock for user(%v) with err = %v", mail.userID, err)
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
		user.UserID = mail.userID
		user.BotPlatform = int(mail.platform)
		user.Stocks = append(user.Stocks, stockID)
		user.LastUpdate = time.Now().UTC()
		_, err = handle.userColl.UpdateOne(context.Background(), bson.M{"user_id": mail.userID}, bson.M{"$set": user}, options.Update().SetUpsert(true))
		if err != nil {
			responseMail.toMsg = fmt.Sprintf("Add stock fail (%v)", err)
		}
	}
	responseMail.toMsg = fmt.Sprintf("Add %v ok", stockID)
	handle.mailbox <- responseMail
}
