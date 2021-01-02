package handlers

import (
	"go.mongodb.org/mongo-driver/mongo"
	tg "gopkg.in/tucnak/telebot.v2"
)

// RegisterHandlers register bot handers
func RegisterHandlers(bot *tg.Bot, collection *mongo.Collection) {
	responseCallback := func(to tg.Recipient, what interface{}, options ...interface{}) {
		bot.Send(to, what, options...)
	}
	bot.Handle(getEchoHandler(responseCallback))
	bot.Handle(getAddStockHandler(responseCallback, collection))
}
