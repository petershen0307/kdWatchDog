package handlers

import (
	"context"
	"log"
	"sort"
	"strings"

	"github.com/petershen0307/kdWatchDog/models"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

const listCommand = "/list"

// ListStock handle list command
func (handle *Handler) ListStock(mail *Mail) {
	defer func() {
		handle.mailbox <- *mail
	}()
	queryUser := models.User{}
	err := handle.userColl.FindOne(
		context.Background(),
		bson.M{"user_id": mail.userID},
	).Decode(&queryUser)
	if err == mongo.ErrNoDocuments {
		return
	}
	if err != nil {
		mail.toMsg = err.Error()
		log.Printf("Query user(%v) with err = %v", mail.userID, err)
		return
	}

	// [ToDo] add stock name and region(tw, us)
	// Region ID Name
	sort.Strings(queryUser.Stocks)
	mail.toMsg = strings.Join(queryUser.Stocks, "\n")
}
